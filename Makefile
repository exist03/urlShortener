generate-gateway:
	protoc --proto_path=api/proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/urlShortener/*proto
psql:
	echo STORAGE_TYPE=psql>.env
	docker compose --profile db up --build
redis:
	echo STORAGE_TYPE=redis>.env
	docker compose --profile db up --build
in-memory:
	echo STORAGE_TYPE=inMemo>.env
	docker compose --profile memory up --build
tests:
	go test -cover -v ./internal/service