package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func (s *Server) Add(c context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Println("On the Server side, function is invoked")
	num1 := in.A
	num2 := in.B

	sum := num1 + num2

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}
