package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"grpc-hello/greet/grpc-hello/greet"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := greet.NewGreetServiceClient(conn)

	stream, err := client.GreetChat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Sending & receiving happens concurrently
	waitc := make(chan struct{})

	// Receiving go-routine
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Printf("Stream closed: %v", err)
				close(waitc)
				return
			}
			log.Printf("Server: %s", res.GetMessage())
		}
	}()

	// Sending go-routine (manual input for demonstration)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter names to greet (type 'exit' to stop):")
	for scanner.Scan() {
		name := scanner.Text()
		if name == "exit" {
			stream.CloseSend()
			break
		}
		if err := stream.Send(&greet.GreetRequest{Name: name}); err != nil {
			log.Fatalf("Send error: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	<-waitc
}
