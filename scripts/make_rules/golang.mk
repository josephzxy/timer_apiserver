GO := go

GO_MAIN_FILE := $(PROJECT_ROOT)/cmd/apiserver.go
GO_BIN := apiserver
GO_MOD_FILE := $(PROJECT_ROOT)/go.mod
GO_MODULE := $(word 2, $(subst $(SPACE), , $(shell cat $(GO_MOD_FILE) | head -n 1)))

GOLANG_MK_PREFIX := "Golang:"

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

.PHONY: go.build
go.build: go.build.$(GO_BIN).$(PLATFORM)

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
