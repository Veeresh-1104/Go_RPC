package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

func DoLongGreet(c pb.GreetServiceClient) {
	log.Println("DoGreetLong is invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Veersh"},
		{FirstName: "Sanath"},
		{FirstName: "Arnav"},
		{FirstName: "Alweeeny"},
		{FirstName: "Yashas"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalln("Error in DoLongGreet client side")
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	//Now we need to receive the response back
	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error while receiving the response, client side")
	}

	log.Printf("The final response \n----- %s\n-----\n", resp.Result)

}
