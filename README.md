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
protoc --go_out=server-golang --go-grpc_out=server-golang proto/hello.proto
protoc --go_out=client-golang --go-grpc_out=client-golang proto/hello.proto
```

## Node
https://www.npmjs.com/package/protoc-gen-grpc
https://medium.com/nerd-for-tech/creating-a-grpc-server-and-client-with-node-js-and-typescript-bb804829fada
https://rsbh.dev/blogs/grpc-with-nodejs-typescript
https://fuchsia.googlesource.com/third_party/protobuf/+/HEAD/js/README.md

```shell
#sudo npm install -g protoc-gen-js
sudo npm install -g protoc-gen-grpc
```

```shell

```

```shell
protoc-gen-grpc --js_out=import_style=commonjs,binary:client-nodejs --grpc_out=grpc_js:client-nodejs proto/hello.proto
protoc-gen-grpc-ts --ts_out=client-nodejs  proto/hello.proto
```

```shell
protoc --plugin=./client-nodejs/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=client-nodejs --ts_proto_opt=outputServices=grpc-js proto/hello.proto
```

# Deselected Variants
Before I discovered [ts-proto](https://github.com/stephenh/ts-proto) for JavaScript code generation,
I tried a number of other approaches found in the official documentation and in articles:

### Dynamic Generation
The [gRPC home page](https://grpc.io/docs/languages/node/basics/) proposes to load the `.proto` files into the
target application, and generate service descriptors and client stub definitions dynamically using
[@grpc/grpc-js](https://www.npmjs.com/package/@grpc/grpc-js) and [@grpc/proto-loader](https://www.npmjs.com/package/@grpc/proto-loader).
This approach is also used in several articles, such as
[How to Build Your First Node.js gRPC API](https://www.trendmicro.com/en_us/devops/22/f/grpc-api-tutorial.html).

```shell
nom install @grpc/grpc-js @grpc/proto-loader
```
((TODO: Code example?))

However, I didn't want to have my skeletons and stubs to be auto-generated,
because I wanted a dedicated code-generation step, clearly separated from actual code usage.
Moreover, I wanted to use the same approach for different languages, such as Go.

### Static Generation with a `protoc` Plugin
The official Protocol Buffer compiler [protoc](https://grpc.io/docs/protoc-installation/) supports the generation of
JavaScript code with flag `--js_out`, given that [protoc-gen-js](https://www.npmjs.com/package/protoc-gen-js) is globally installed:
```shell
npm install -g protoc-gen-js
```
Then the Protocol Buffers can be generated as follows:
```shell
protoc --js_out=import_style=commonjs,binary:client-nodejs proto/hello.proto
```
The additional generation of gRPC skeletons and stubs requires the use of a `protoc` plugin,
which comes as part of the [grpc-tools](https://www.npmjs.com/package/grpc-tools):
```shell
npm install -g grpc-tools
```
Then, the complete compiler call is a follows:
```shell
protoc --js_out=import_style=commonjs,binary:client-nodejs --grpc_out=client-nodejs --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` proto/hello.proto
```
The plugin syntax is described in article [Creating a gRPC server and client with Node.js and TypeScript](https://medium.com/nerd-for-tech/creating-a-grpc-server-and-client-with-node-js-and-typescript-bb804829fada),
which unfortunately doesn't mention the need for global instalations of `protoc-gen-js` and `grpc-tools`.

### Static Generation with `protoc-gen-grpc`
The need for three global installations (`protoc`, `protoc-gen-js`, and `grpc-tools`)
plus the verbose argument syntax of `protoc` makes the previous variant rather complex.
An alternative is [protoc-gen-grpc](https://www.npmjs.com/package/protoc-gen-grpc),
which ships with `protoc` and all required plugins:
```shell
npm install -g protoc-gen-grpc
```
The argument syntax is simplified compared to the direct call to `protoc`(no need to provide a `--plugin` option): 
```shell
protoc-gen-grpc --js_out=import_style=commonjs,binary:client-nodejs --grpc_out=grpc_js:client-nodejs proto/hello.proto
```
Under the hood, `protoc-gen-grpc` delegates its arguments to `protoc`and passes the proper plugin.

As a nice add-on, the package can also generate TypeScript signatures:
```shell
protoc-gen-grpc-ts --ts_out=client-nodejs  proto/hello.proto
```

### Static Generation with `ts-proto`
All variants discussed so far use the built-in JavaScript generation of `protoc`.
There are different options for the output format. CommonJS appears the most modern one,
because, as the [project page](https://github.com/protocolbuffers/protobuf-javascript) states
> Support for ES6-style imports is not implemented yet.

This leads to generated code that is hard to read for users of modern JavaScript.
In addition, the use of `require` makes it difficult to integrate the generated code into
[ECMAScript modules](https://nodejs.org/api/esm.html), i.e. projects specified as
`"type": "module"` in the `package.json`.

Project [ts-proto](https://github.com/stephenh/ts-proto) goes a radically different way
and replaces the built-in JavaScript generation of `protoc` by a TypeScript generator.
It still uses [protobufjs](https://www.npmjs.com/package/protobufjs) internally, apparently for reading `.proto` files.
However, the generated code is quite different from the variants discussed so far, with a different API
(for example, there are no getters and setters, just fields).
The differences are a great opportunity for easier integration into your target projects,
but there is also a risk in becoming dependent on `ts-proto`.

Only one package needs to be installed:
```shell
npm install ts-code
```
The generator is passed as plugin to `protoc`.
It provides a wealth of extra options, including one that generates gRPC service definitions and stubs:

```shell
protoc --plugin=./client-nodejs/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=client-nodejs --ts_proto_opt=outputServices=grpc-js proto/hello.proto
```
All code is written to one idiomatic Typescript file.

A toy gRPC client can be as simple as
```typescript
import * as grpc from '@grpc/grpc-js'
import { GreeterClient, HelloRequest } from './proto/hello'

const client = new GreeterClient('localhost:5005', grpc.credentials.createInsecure())
const request = HelloRequest.create({ name: 'Hans' })
client.sayHello(request, (error, response) => {
    console.log(error ? error : response.message)
})
```
Unfortunately, the generated functions use callbacks.
A promise-based API would be more elegant and would fit better with the clean Typescript approach.

Many more details of creating gRCP server and client applications with `ts-proto` are provided
by article [NodeJS Microservice with gRPC and TypeScript](https://rsbh.dev/blogs/grpc-with-nodejs-typescript).