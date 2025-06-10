package main

import (
	"context"
	"grpc-hello/greet/grpc-hello/greet"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greetServer struct {
	greet.UnimplementedGreetServiceServer
}

func (s *greetServer) SayHello(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	log.Printf("Received request from: %s", req.GetName())
	message := "Hello " + req.GetName()
	return &greet.GreetResponse{Message: message}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	greet.RegisterGreetServiceServer(grpcServer, &greetServer{})

	log.Println("gRPC Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
