GRPC_MK_PREFIX := "gRPC:"

## grpc.protoc: Parse proto files and generate output
.PHONY: grpc.protoc
grpc.protoc: tools.verify.protoc-gen-go tools.verify.protoc-gen-go-grpc
	@echo "=======> $(GRPC_MK_PREFIX) running protoc"
	@protoc --go_opt=paths=source_relative --go_out=. \
	--go-grpc_opt=paths=source_relative --go-grpc_out=. \
	./api/grpc/timer.proto
