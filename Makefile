.PHONY: grpc-server-compile
grpc-server-compile: 
	protoc --go_out=./pkg/proto/server --go-grpc_out=./pkg/proto/server ./pkg/proto/server/server.proto

test:
	go test -v -timeout=1m ./...
