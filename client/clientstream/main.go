package main

import (
	"context"
	"grpc-hello/greet/grpc-hello/greet"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)

	}
	defer conn.Close()
	client := greet.NewGreetServiceClient(conn)
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error calling GreetEveryone: %v", err)

	}
	names := []string{"Anish", "Jhon", "Bob", "Charlie"}
	for _, name := range names {
		log.Printf("Sending name: %s", name)
		if err := stream.Send(&greet.GreetRequest{Name: name}); err != nil {
			log.Fatalf("error sending name: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving reposnse: %v", err)
	}
	log.Printf("Final response: %s", res.GetMessage())
}
