SERVER_PROTO_DIR=./internal/presentation/grpc/protobuf

# Generate all protobuf source codes
proto-generate-server:
	@echo "Generating protobuf source code ..."
	protoc --proto_path=$(SERVER_PROTO_DIR) \
		   --go_out=$(SERVER_PROTO_DIR) \
		   --go-grpc_out=$(SERVER_PROTO_DIR) \
		   $(wildcard $(SERVER_PROTO_DIR)/*.proto)
	@echo "Done"