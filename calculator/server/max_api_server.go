package main

import (
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

var CurrMax int32 = -1e6

func (s *Server) MaxApiBiStream(stream pb.AddTwoNumbrs_MaxApiBiStreamServer) error {

	for {

		num, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln("Error receiving the number")
		}

		CurrMax = max(CurrMax, num.Num)

		err = stream.Send(&pb.MaxApiResponse{
			Max: CurrMax,
		})

		if err != nil {
			log.Fatalln("Error while Sending the stream to the Client")
		}

	}

}
