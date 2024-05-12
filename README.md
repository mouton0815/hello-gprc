# Playing with gRPC





## Go
https://grpc.io/docs/languages/go/quickstart/

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```
```shell
protoc --go_out=server-golang --go-grpc_out=server-golang proto/hello.proto
protoc --go_out=client-golang --go-grpc_out=client-golang proto/hello.proto
```

## Node
https://www.npmjs.com/package/protoc-gen-grpc
https://medium.com/nerd-for-tech/creating-a-grpc-server-and-client-with-node-js-and-typescript-bb804829fada

```shell
#sudo npm install -g protoc-gen-js
sudo npm install -g protoc-gen-grpc
```

```shell
protoc-gen-grpc --js_out=import_style=commonjs,binary:client-nodejs --grpc_out=grpc_js:client-nodejs proto/hello.proto
protoc-gen-grpc-ts --ts_out=client-nodejs  proto/hello.proto
```
