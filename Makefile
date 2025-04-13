proto:
	protoc \
		--proto_path=proto \
		--go_out=proto \
		--go-grpc_out=proto \
		proto/*.proto
	@echo "Proto files generated in the 'proto' directory."