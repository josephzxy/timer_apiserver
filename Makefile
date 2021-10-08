# Build all by default
.DEFAULT_GOAL := all

MAIN_PREFIX := "Main:"

## all: Run the default build procedure(tidy format lint test build)
.PHONY: all
all: tidy format test lint build

# Includes
include scripts/make_rules/common.mk # must always be the first
include scripts/make_rules/golang.mk
include scripts/make_rules/tools.mk
include scripts/make_rules/docker.mk
include scripts/make_rules/swagger.mk
include scripts/make_rules/mysql.mk
include scripts/make_rules/grpc.mk

# Usages
define USAGE_OPTIONS

Options:
PLATFORM	The target platform of the binary to build. In the format "OS_ARCH". Default to "$${GOOS}_$${GOARCH}"
DKR_ARCH	The target architecture of the docker image to build. Default to ARCH in Option PLATFORM
endef
export USAGE_OPTIONS

# Targets

## tidy: Tidy dependencies
.PHONY: tidy
tidy: go.tidy

## format: Format source code
.PHONY: format
format: go.format

## lint: Lint source code
.PHONY: lint
lint: go.lint

## test: Run unit tests
.PHONY: test
test: go.test

## build: Build binary from source code
.PHONY: build
build: go.build

## gen: Generate necessary files(e.g. gRPC stub, gomock files)
.PHONY: gen
gen: 
	@echo "=======> $(MAIN_PREFIX) generating files"
	@$(MAKE) go.mock
	@$(MAKE) grpc.protoc

## docker: Build docker image
.PHONY: docker
docker: docker.build

## swagger: Generate swagger doc for RESTful API and launch a local HTTP server to display
.PHONY: swagger
swagger: swagger.generate swagger.serve

## clean: Remove temporary files generated during building
.PHONY: clean
clean:
	@echo "=======> $(MAIN_PREFIX) cleaning"
	@rm -rvf $(OUTPUT_DIR)
	@$(MAKE) go.clean

## help: print the help message and exit
.PHONY: help
help: Makefile
	@echo -e "Usage: make <TARGETS> <OPTIONS>... \n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
