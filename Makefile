.PHONY: grpc-server-compile

grpc-server-compile: 
	protoc --go_out=./pkg/proto/server --go-grpc_out=./pkg/proto/server ./pkg/proto/server/server.proto

