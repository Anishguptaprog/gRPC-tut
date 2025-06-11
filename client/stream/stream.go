package main

import (
	"bufio"
	"context"
	"fmt"
	"grpc-hello/greet/grpc-hello/greet"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)

	}
	defer conn.Close()
	client := greet.NewGreetServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	name = name[:len(name)-1] // Remove the newline character
	req := &greet.GreetRequest{
		Name: name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("Error in calling GreetManyTimes: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		log.Printf("Response: %s", res.GetMessage())
	}

}
