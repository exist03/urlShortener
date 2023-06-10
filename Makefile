generate-gateway:
	protoc --proto_path=api/proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/ozonShortLinks/*proto

