dc:
	docker compose --verbose up --build --remove-orphans

run:
	GRPC_HOST=0.0.0.0 \
	GRPC_PORT=9000 \
	HTTP_HOST=0.0.0.0 \
	HTTP_PORT=9090 \
	IMMUDB_HOST=localhost \
	IMMUDB_PORT=3322 \
	IMMUDB_DATABASE=defaultdb \
	IMMUDB_USERNAME=immudb \
	IMMUDB_PASSWORD=immudb \
	go run --race cmd/server/main.go

install-go-tools: install-go-lib-tools
	go install github.com/yoheimuta/protolint/cmd/protolint@v0.42.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

install-go-lib-tools:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.29


protos:
	cd proto && make proto


