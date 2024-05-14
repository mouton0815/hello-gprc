import { credentials } from '@grpc/grpc-js'
import { GreeterClient, HelloRequest } from './proto/hello'

const client = new GreeterClient('localhost:5005', credentials.createInsecure())
const request = HelloRequest.create({ name: 'Hans' })
client.sayHello(request, (error, response) => {
    console.log(error ? error : response.message)
})