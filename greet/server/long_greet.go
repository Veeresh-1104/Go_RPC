package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("Long Greet invoked on server with \n")

	res := ""
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: fmt.Sprintf("Reached the END of the request Stream\nReceived Stream : %v", res),
			})
		}

		if err != nil {
			log.Fatalln("Error on Long Greet - server")
		}

		log.Printf("Processing data with : %v\n", msg.FirstName)

		res += fmt.Sprintf("Hello %s\n", msg.FirstName)

	}
}
