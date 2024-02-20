package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func PrimeDecompClient(c pb.AddTwoNumbrsClient) {
	log.Println("Function invoked on the client side")

	req := pb.PrimeRequest{
		Number: 32,
	}

	res, err := c.PrimeServerStream(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error while receiving the response %v\n", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			log.Fatalln("Error on client side, couldn't read the stream")
		}

		log.Println("The Decomposition factors : ", msg.Answer)

	}
}
