proto:
	protoc \
		--proto_path=proto \
		--go_out=generated \
		--go-grpc_out=generated \
		proto/*.proto
	@echo "Proto files generated in the 'generated' directory."