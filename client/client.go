package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	unaryResponse, err := client.UnaryCall(context.Background(), &pb.RequestMessage{Message: "Hello Unary", RequestedAt: timestamppb.Now()})
	respTime1 := time.Since(start1)
	log.Printf("Unary Response Time: %v", respTime1)
	if err != nil {
		log.Fatalf("UnaryCall error: %v", err)
	}
	fmt.Println("Unary Response:", unaryResponse.Message)

	// Server-to-Client Streaming Call
	stream, err := client.ServerToClientStreamingCall(context.Background(), &pb.RequestMessage{Message: "Hello Stream", RequestedAt: timestamppb.Now()})
	if err != nil {
		log.Fatalf("ServerToClientStreamingCall error: %v", err)
	}

	for {
		response, err := stream.Recv()
		delay := time.Since(response.RespondedAt.AsTime())
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		fmt.Printf("Stream Response: %s, delay: %v", response.Message, delay)
	}

	log.Println("Client completed")
}
