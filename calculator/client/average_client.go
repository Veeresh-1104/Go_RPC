package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func AverageClient(c pb.AddTwoNumbrsClient) {

	log.Println("Average Client invoked - client side")

	reqs := []*pb.AverageRequest{
		{Num: 12.3},
		{Num: 43.69},
		{Num: 31.24},
		{Num: 12.34},
	}

	clientStream, err := c.AverageClientStream(context.Background())

	if err != nil {
		log.Fatalln("Error in context creation")
	}

	for _, req := range reqs {

		clientStream.Send(req)
		time.Sleep(2 * time.Second)
	}

	resp, err := clientStream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Couldn't receive the final response")
	}

	log.Printf("The Average Received is : %f\n", resp.Avg)
}
