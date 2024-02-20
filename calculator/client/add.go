package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func AddNum(c pb.AddTwoNumbrsClient) {

	log.Println("Add Num function is invoked")

	res, err := c.Add(context.Background(), &pb.AddRequest{
		A: 10.3,
		B: 23.7,
	})

	if err != nil {
		log.Fatalf("The error : %v", err)
		log.Fatalln("Error In Add Num, client side")
	}

	log.Printf("Response received : %v\n", res.Sum)

}
