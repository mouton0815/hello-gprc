package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/main/internal"
	"log"
	"net"
)

type GreeterServerImpl struct {
	internal.UnimplementedGreeterServer
}

func (GreeterServerImpl) SayHello(_ context.Context, req *internal.HelloRequest) (*internal.HelloReply, error) {
	return &internal.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5005))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	internal.RegisterGreeterServer(s, &GreeterServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
