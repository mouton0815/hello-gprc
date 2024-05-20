package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/main/internal"
	"log"
	"math/rand"
	"net"
)

var NAMES = []string{
	"Hans",
	"Inge",
	"Kurt",
	"Anna",
	"Fred",
	"Klara",
	"Klaus",
	"Olaf",
	"Lara",
	"Lars",
}

type GreeterServerImpl struct {
}

func (GreeterServerImpl) SayHello(_ context.Context, req *internal.HelloRequest) (*internal.HelloReply, error) {
	return &internal.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func (GreeterServerImpl) GetNames(req *internal.NameRequest, server internal.Greeter_GetNamesServer) error {
	log.Printf("Start streaming %d names", req.Count)
	for i := 0; i < int(req.Count); i++ {
		name := NAMES[rand.Intn(len(NAMES))]
		reply := internal.NameReply{Name: name}
		if err := server.Send(&reply); err != nil {
			log.Println("error generating response")
			return err
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
	internal.RegisterGreeterServer(s, &GreeterServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
