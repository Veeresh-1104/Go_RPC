package main

import (
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func (s *Server) PrimeServerStream(in *pb.PrimeRequest, stream pb.AddTwoNumbrs_PrimeServerStreamServer) error {

	log.Printf("The Number received is :%d", in.Number)
	num := in.Number

	var fact int32 = 2

	for num > 1 {
		if num%fact == 0 {
			num = num / fact
			stream.Send(&pb.PrimeResponse{
				Answer: fact,
			})
		} else {
			fact++
		}
	}

	return nil
}
