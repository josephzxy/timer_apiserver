.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.install.goimports
tools.install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest
