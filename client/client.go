package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/aakashbajaj/istio-grpc-test/internal"

	"google.golang.org/grpc"
)

func main() {
	connStart := time.Now()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	connTime := time.Since(connStart)
	log.Printf("Connection Response Time: %v", connTime)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)

	// Unary Call
	start1 := time.Now()
	unaryResponse, err := client.UnaryCall(context.Background(), &pb.RequestMessage{Message: "Hello Unary"})
	respTime1 := time.Since(start1)
	log.Printf("Unary Response Time: %v", respTime1)
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
