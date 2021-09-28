# Build all by default
.DEFAULT_GOAL := all

MAIN_PREFIX := "Main:"

.PHONY: all
all: tidy format lint test build

# Includes
include scripts/make_rules/common.mk # must always be the first
include scripts/make_rules/golang.mk
include scripts/make_rules/tools.mk
include scripts/make_rules/docker.mk

.PHONY: tidy
tidy: go.tidy

.PHONY: format
format: go.format

.PHONY: lint
lint: go.lint

.PHONY: test
test: go.test

.PHONY: build
build: go.build

.PHONY: docker
docker: docker.build

.PHONY: clean
clean:
	@echo "=======> $(MAIN_PREFIX) cleaning"
	@rm -rvf $(OUTPUT_DIR)
	@$(MAKE) go.clean
