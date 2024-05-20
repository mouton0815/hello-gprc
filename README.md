# Playing with gRPC


# Installation

You need `protoc`, the Protocol Buffer compiler.
Please follow the instructions on https://grpc.io/docs/protoc-installation/.



## Go
https://grpc.io/docs/languages/go/quickstart/

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```
```shell
protoc --go_out=server --go-grpc_out=require_unimplemented_servers=false:server hello.proto
protoc --go_out=client --go-grpc_out=client hello.proto
```
