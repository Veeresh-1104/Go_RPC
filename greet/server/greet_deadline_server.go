package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("deadline invoked \n")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Deadline exceeded")
			return nil, status.Errorf(codes.Canceled, "Deadline exceeded")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
