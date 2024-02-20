package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DoGreetDeadline(c pb.GreetServiceClient, timeout time.Duration) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Hola",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		fmt.Printf("The error : %v\n", err)
		if ok {
			if e.Code() == codes.Canceled {
				log.Printf("Deadline Exceeded \n")
			} else {
				log.Fatalf("heyyo, its unexpected \n")
				return
			}
		} else {
			log.Fatalf("Non gRPC error : %v", e)
		}

	}

	log.Printf("Dealdine : %s \n", res.Result)

}
