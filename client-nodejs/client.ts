import parseArgs from 'minimist'
import { credentials, ServiceError } from '@grpc/grpc-js'
import { OrderRequest, OrderResponse, OrdersClient, StatsResponse } from './orders'

// Get command line arguments
const args = parseArgs(process.argv.slice(2))
const minDate = args['mindate'] || 0
const maxDate = args['maxdate'] || Math.pow(2, 31) - 1

const addr = 'localhost:5005'
const client = new OrdersClient(addr, credentials.createInsecure())

// First retrieve the order stats
client.getStats({}, (error : ServiceError, response : StatsResponse) => {
    console.log(error ? error.message : `#Orders: ${response.count}`)

    // Next fetch the orders stream from the server
    const request = OrderRequest.create({ minDate, maxDate })
    const orderStream = client.getOrders(request)
    orderStream.on('data', (order: OrderResponse) => {
        console.log(order)
    })
})
