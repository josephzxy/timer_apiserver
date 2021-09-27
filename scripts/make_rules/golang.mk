GO := go

GOLANG_MK_PREFIX := "Golang:"

.PHONY: go.tidy
go.tidy:
	@echo "=======> $(GOLANG_MK_PREFIX) tidying dependencies ..."
	@$(GO) mod tidy

.PHONY: go.format
go.format:
	@echo "=======> $(GOLANG_MK_PREFIX) formatting source code ..."
	@gofmt -s -w .

.PHONY: go.clean
go.clean:
	@echo "=======> $(GOLANG_MK_PREFIX) cleaning ..."
	@go clean -x
