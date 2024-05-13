import * as grpc from '@grpc/grpc-js'
import { GreeterClient, HelloRequest } from './proto/hello'

const client = new GreeterClient('localhost:5005', grpc.credentials.createInsecure())
const request = HelloRequest.create()
request.name = 'Hans'
client.sayHello(request, (error, response) => {
    if (error) {
        console.error(error);
    } else {
        console.log(response.message);
    }
})