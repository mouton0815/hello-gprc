import parseArgs from 'minimist'
import { credentials } from '@grpc/grpc-js'
import { OrderRequest, OrderResponse, OrdersClient } from './orders'

// Get command line arguments
const args = parseArgs(process.argv.slice(2))
const minDate = args['mindate'] || 0
const maxDate = args['maxdate'] || Math.pow(2, 31) - 1

const client = new OrdersClient('localhost:5005', credentials.createInsecure())
const request = OrderRequest.create({ minDate, maxDate })

// Fetch the orders stream from the server
const orderStream = client.getOrders(request)
orderStream.on('data', (order: OrderResponse) => {
    console.log(order)
})
