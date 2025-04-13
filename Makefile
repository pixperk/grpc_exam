proto:
	protoc \
		--proto_path=proto \
		--go_out=proto \
		--go-grpc_out=proto \
		proto/*.proto
	@echo "Proto files generated in the 'proto' directory."

server:
	go run server/main.go
	
client_unary:
	go run client/main.go unary
	
client_server:
	go run client/main.go server
	
client_client:
	go run client/main.go client
	
client_bidi:
	go run client/main.go bidi
	
.PHONY: proto server client_unary client_server client_client client_bidi
