package main

import (
	"context"
	"fmt"
	"grpc-hello/greet/grpc-hello/greet"
	"io"
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
func (s *greetServer) GreetEveryone(stream greet.GreetService_GreetEveryoneServer) error {
	log.Println("Receiving stream of greetings")
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			message := fmt.Sprintf("Hello %s", joinNames(names))
			return stream.SendAndClose(&greet.GreetResponse{Message: message})
		}
		if err != nil {
			return err
		}
		names = append(names, req.GetName())
	}

}
func joinNames(names []string) string {
	return joinWithComma(names)
}
func joinWithComma(names []string) string {
	result := ""
	for i, name := range names {
		if i > 0 {
			result += ", "
		}
		result += name
	}
	return result
}
func (s *greetServer) GreetChat(stream greet.GreetService_GreetChatServer) error {
	log.Println("starting bidi stream...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("bidi stream closed by client")
			return nil
		}
		if err != nil {
			return err
		}
		resp := &greet.GreetResponse{
			Message: "Hello " + req.GetName(),
		}
		log.Printf("sending greeting for: %s", req.GetName())
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
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
