generate-gateway:
	protoc --proto_path=api/proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/urlShortener/*proto
psql:
	echo STORAGE_TYPE=psql>.env
	docker compose up --build
redis:
	echo STORAGE_TYPE=redis>.env
	docker compose up --build
