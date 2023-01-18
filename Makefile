depend:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

build:
	buf generate --path api/**/

all: depend build