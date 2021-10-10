.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.install.goimports
tools.install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools.install.gci
tools.install.goimports:
	@$(GO) install github.com/daixiang0/gci@latest

.PHONY: tools.install.golangci-lint
tools.install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: tools.install.swagger
tools.install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: tools.install.go-migrate
tools.install.go-migrate:
	@$(GO) install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: tools.install.mockgen
tools.install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: tools.install.protoc-gen-go
tools.install.protoc-gen-go:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: tools.install.protoc-gen-go-grpc
tools.install.protoc-gen-go-grpc:
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
