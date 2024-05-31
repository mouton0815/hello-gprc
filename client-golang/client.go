package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_go/main/proto"
	"io"
	"log"
	"math"
	"time"
)

func main() {
	// Get command line arguments
	var minDate = flag.Int("mindate", 0, "minimum order date")
	var maxDate = flag.Int("maxdate", math.MaxInt32, "maximum order date")
	flag.Parse()

	conn, err := grpc.Dial("localhost:5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	client := proto.NewOrdersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Fetch the orders stream from the server
	stream, err := client.GetOrders(ctx, &proto.OrderRequest{MinDate: int32(*minDate), MaxDate: int32(*maxDate)})
	if err != nil {
		panic(fmt.Sprintf("Could not stream: %v", err))
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("Stream interrupted: %v", err))
		}
		log.Printf("{%v}", order)
	}
}
