import { sendUnaryData, Server, ServerCredentials, ServerUnaryCall, ServerWritableStream } from '@grpc/grpc-js'
import { OrderRequest, OrderResponse, OrdersService, StatsRequest, StatsResponse } from './orders'
import { ORDERS } from './order-data'

function getStats(call: ServerUnaryCall<StatsRequest, StatsResponse>, callback: sendUnaryData<StatsResponse>) {
    callback(null, StatsResponse.create({ count: ORDERS.length }))
}

function getOrders(call: ServerWritableStream<OrderRequest, OrderResponse>) {
    ORDERS.forEach(order => {
        call.write(order)
    })
    call.end()
}

const server = new Server()
server.addService(OrdersService, { getStats, getOrders })
server.bindAsync('0.0.0.0:5005', ServerCredentials.createInsecure(), (error, port) => {
    console.log(error ? error.message : 'Server listening on port ' + port)
})
