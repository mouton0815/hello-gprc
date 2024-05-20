package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_go/main/proto"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	client := proto.NewGreeterClient(conn)

	// First say hello
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	hello, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Client"})
	if err != nil {
		panic(fmt.Sprintf("Could not greet: %v", err))
	}
	log.Println(hello.GetMessage())

	// Now fetch the names stream from the server
	stream, err := client.GetNames(ctx, &proto.NameRequest{Count: 7})
	if err != nil {
		panic(fmt.Sprintf("Could not stream: %v", err))
	}
	for {
		name, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("Stream interrupted: %v", err))
		}
		log.Printf("Hello %s", name.Name)
	}
}