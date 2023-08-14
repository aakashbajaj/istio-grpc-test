package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "your-module-name/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) UnaryCall(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	message := fmt.Sprintf("Received: %s", req.Message)
	return &pb.ResponseMessage{Message: message}, nil
}

func (s *server) ServerToClientStreamingCall(req *pb.RequestMessage, stream pb.SampleService_ServerToClientStreamingCallServer) error {
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Response %d: %s", i+1, req.Message)
		err := stream.Send(&pb.ResponseMessage{Message: message})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterSampleServiceServer(srv, &server{})
	log.Println("Server started on :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
