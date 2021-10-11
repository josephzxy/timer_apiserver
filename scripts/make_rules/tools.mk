## tools.verify.%: Verify if the tool of the given name is installed. Install it if not
.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

## tools.install.goimports: Install goimports
.PHONY: tools.install.goimports
tools.install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

## tools.install.gci: Install gci
.PHONY: tools.install.gci
tools.install.gci:
	@$(GO) install github.com/daixiang0/gci@latest

## tools.install.gofumpt: Install gofumpt
.PHONY: tools.install.gofumpt
tools.install.gofumpt:
	@$(GO) install mvdan.cc/gofumpt@latest

## tools.install.golangci-lint: Install golangci-lint
.PHONY: tools.install.golangci-lint
tools.install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

## tools.install.swagger: Install swagger
.PHONY: tools.install.swagger
tools.install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

## tools.install.go-migrate: Install go-migrate
.PHONY: tools.install.go-migrate
tools.install.go-migrate:
	@$(GO) install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## tools.install.mockgen: Install mockgen
.PHONY: tools.install.mockgen
tools.install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

## tools.install.protoc-gen-go: Install protoc-gen-go
.PHONY: tools.install.protoc-gen-go
tools.install.protoc-gen-go:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## tools.install.protoc-gen-go-grpc: Install protoc-gen-go-grpc
.PHONY: tools.install.protoc-gen-go-grpc
tools.install.protoc-gen-go-grpc:
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

