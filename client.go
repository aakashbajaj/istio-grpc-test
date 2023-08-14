package main

import (
	"context"
	"fmt"
	"log"

	pb "your-module-name/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)

	// Unary Call
	unaryResponse, err := client.UnaryCall(context.Background(), &pb.RequestMessage{Message: "Hello Unary"})
	if err != nil {
		log.Fatalf("UnaryCall error: %v", err)
	}
	fmt.Println("Unary Response:", unaryResponse.Message)

	// Server-to-Client Streaming Call
	stream, err := client.ServerToClientStreamingCall(context.Background(), &pb.RequestMessage{Message: "Hello Stream"})
	if err != nil {
		log.Fatalf("ServerToClientStreamingCall error: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println("Stream Response:", response.Message)
	}

	log.Println("Client completed")
}
