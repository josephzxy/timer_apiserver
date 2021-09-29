GO := go

GO_MAIN_FILE := $(PROJECT_ROOT)/cmd/apiserver.go
GO_BIN := apiserver
GO_MOD_FILE := $(PROJECT_ROOT)/go.mod
GO_MODULE := $(word 2, $(subst $(SPACE), , $(shell cat $(GO_MOD_FILE) | head -n 1)))

GO_SUPPORTED_VERSIONS := 1.17

GOLANG_MK_PREFIX := "Golang:"

NO_TEST_PKGS := "\
github.com/josephzxy/timer_apiserver/cmd|\
github.com/josephzxy/timer_apiserver/api/rest/swagger/docs|\
github.com/josephzxy/timer_apiserver/internal/resource/model/v1|\
github.com/josephzxy/timer_apiserver/internal/pkg/log\
"
NO_TEST_PKGS := $(shell echo '$(NO_TEST_PKGS)' | tr -d '[:space:]')

.PHONY: go.tidy
go.tidy:
	@echo "=======> $(GOLANG_MK_PREFIX) tidying dependencies"
	@$(GO) mod tidy

.PHONY: go.format
go.format: tools.verify.goimports
	@echo "=======> $(GOLANG_MK_PREFIX) formatting source code"
	@gofmt -s -w $(PROJECT_ROOT)
	@goimports -w -local $(GO_MODULE) $(PROJECT_ROOT)

.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@echo "=======> $(GOLANG_MK_PREFIX) linting source code"
	@golangci-lint run $(PROJECT_ROOT)/...

.PHONY: go.test
go.test:
	@echo "=======> $(GOLANG_MK_PREFIX) running unit tests"
	@mkdir -p $(OUTPUT_DIR)
	@set -o pipefail; $(GO) test -v -short -race -timeout=10m \
	-coverprofile=$(OUTPUT_DIR)/coverage.out \
	`$(GO) list $(PROJECT_ROOT)/... | egrep -v $(NO_TEST_PKGS)`
	@$(GO) tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html
	@$(GO) tool cover -func=$(OUTPUT_DIR)/coverage.out

.PHONY: go.build.verify
go.build.verify:
ifeq ($(shell $(GO) version | egrep -q '\bgo($(GO_SUPPORTED_VERSIONS))\b' && echo 0 || echo 1), 1)
	$(error Go version not supported. Expecting one of the following: $(GO_SUPPORTED_VERSIONS), got: $(shell $(GO) version))
endif

.PHONY: go.build
go.build: go.build.verify go.build.$(GO_BIN).$(PLATFORM)

.PHONY: go.build.%
go.build.%:
	$(eval BIN := $(word 1, $(subst ., , $*)))
	$(eval PLAT := $(word 2, $(subst ., , $*)))
	$(eval OS := $(word 1, $(subst _, , $(PLAT))))
	$(eval ARCH := $(word 2, $(subst _, , $(PLAT))))
	@echo "=======> $(GOLANG_MK_PREFIX) building binary $(BIN) for $(OS) $(ARCH)"
	$(eval BIN_DIR := $(OUTPUT_DIR)/build/$(OS)/$(ARCH))
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -o $(BIN_DIR)/$(BIN) $(GO_MAIN_FILE)

.PHONY: go.clean
go.clean:
	@echo "=======> $(GOLANG_MK_PREFIX) cleaning"
	@go clean -x `go list $(PROJECT_ROOT)/...`
