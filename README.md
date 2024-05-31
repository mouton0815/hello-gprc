# Playing with gRPC
A simple gRPC server and client that demonstrates
* A basic request-response pattern (method `GetStats`)
* A streaming response triggered by a unary request (method `GetOrders`)

The streaming variant is a modification of the example used in the article
[Implementing gRPC with Golang](https://medium.com/@josueparra2892/implementing-grpc-with-golang-71bd72a4561).

## Installation

You need
* `Go` (https://go.dev/doc/install)
* `protoc` (https://grpc.io/docs/protoc-installation)

To enable `protoc` to generate Go code, you need two additional plugins:
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

See https://grpc.io/docs/languages/go/quickstart/ for more information.

## Code Generation

```shell
protoc --go_out=server-golang --go-grpc_out=require_unimplemented_servers=false:server-golang orders.proto
protoc --go_out=client-golang --go-grpc_out=client-golang orders.proto
```

## Running
Run the server:
```shell
cd server-golang
go run server.go orders.go
```

Run the client (in another shell):
```shell
cd client-golang
go run client.go
```
The client accepts a minimum and maximum value for order dates
```shell
go run client.go --mindate 200 --maxdate 300
```
