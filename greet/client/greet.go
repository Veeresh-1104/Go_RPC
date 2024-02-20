package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

func DoGreet(c pb.GreetServiceClient) {
	log.Println("Do Greet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Veeresh",
	})

	if err != nil {
		log.Fatalf("Error occured in Do greet  %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
