.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.install.goimports
tools.install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools.install.golangci-lint
tools.install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
