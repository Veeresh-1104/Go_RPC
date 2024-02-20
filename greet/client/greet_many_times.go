package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

func GreetManyTimes(c pb.GreetServiceClient) {
	log.Println("GreetManyTimes called on the client side")

	req := &pb.GreetRequest{
		FirstName: "Veeresh",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while receiving the response %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			log.Fatalf("Error, client side \n")
		}

		log.Printf("GreetManyTimes : %v\n", msg.Result)

	}
}
