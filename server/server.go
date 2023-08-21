package server

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"

	pb "github.com/aakashbajaj/istio-grpc-test/internal"

	"google.golang.org/grpc"
)

type server struct {
	ch chan struct{}
	pb.UnimplementedSampleServiceServer
}

func (s *server) UnaryCall(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	message := fmt.Sprintf("Received: %s", req.Message)
	return &pb.ResponseMessage{Message: message}, nil
}

func (s *server) ServerToClientStreamingCall(req *pb.RequestMessage, stream pb.SampleService_ServerToClientStreamingCallServer) error {
	i := 0
	for {
		select {
		case <-s.ch:
			fmt.Printf("terminating application...")
			return nil
		default:
			message := fmt.Sprintf("Response %d: %s", i+1, req.Message)
			err := stream.Send(&pb.ResponseMessage{Message: message, RespondedAt: timestamppb.Now()})
			if err != nil {
				return err
			}
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Start() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterSampleServiceServer(srv, &server{
		ch: make(chan struct{}),
	})
	log.Println("Server started on :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
