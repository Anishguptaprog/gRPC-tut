package main

import (
	"context"
	"fmt"
	"grpc-hello/greet/grpc-hello/greet"
	"log"
	"net"
	"time"

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
func (s *greetServer) GreetManyTimes(req *greet.GreetRequest, stream greet.GreetService_GreetManyTimesServer) error {
	name := req.GetName()
	log.Printf("Streaming greetings to: %s", name)
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Hello %s #%d", name, i)
		res := &greet.GreetResponse{
			Message: msg}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)

	}
	return nil
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
