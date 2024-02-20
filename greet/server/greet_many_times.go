package main

import (
	"fmt"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes invoked on the server side with Request : %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, value %d", in.FirstName, i+1)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
