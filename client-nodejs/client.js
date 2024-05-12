const grpc = require('@grpc/grpc-js')
const { GreeterClient} = require('./proto/hello_grpc_pb')
const { HelloRequest } = require('./proto/hello_pb')


const client = new GreeterClient('localhost:5005', grpc.credentials.createInsecure())
const request = new HelloRequest()
request.setName('Hans')
client.sayHello(request, (error, response) => {
    if (error) {
        console.error(error);
    } else {
        console.log(response.getMessage());
    }
})