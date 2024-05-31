package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/main/proto"
	"log"
	"net"
)

type OrdersServerImpl struct {
}

func (OrdersServerImpl) GetStats(_ context.Context, _ *proto.StatsRequest) (*proto.StatsResponse, error) {
	log.Printf("Return order stats")
	return &proto.StatsResponse{Count: uint32(len(ORDERS))}, nil
}

func (OrdersServerImpl) GetOrders(req *proto.OrderRequest, server proto.Orders_GetOrdersServer) error {
	log.Printf("Start streaming orders from %d to %d", req.MinDate, req.MaxDate)
	for i := 0; i < len(ORDERS); i++ {
		var order = &ORDERS[i]
		if order.Date >= req.MinDate && order.Date <= req.MaxDate {
			if err := server.Send(order); err != nil {
				log.Println("error generating response")
				return err
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5005))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterOrdersServer(s, &OrdersServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
