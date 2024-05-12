# Playing with gRPC

```shell
protoc --go_out=server-golang --go-grpc_out=server-golang proto/hello.proto
protoc --go_out=client-golang --go-grpc_out=client-golang proto/hello.proto
```
