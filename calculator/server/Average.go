package main

import (
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func (s *Server) AverageClientStream(stream pb.AddTwoNumbrs_AverageClientStreamServer) error {
	log.Println("Average server invoked")

	var ans float32 = 0
	var count int = 0

	for {
		num, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Avg: (ans / float32(count)),
			})
		}

		if err != nil {
			log.Fatalln("Error in stream server Average Function")
		}

		count++
		ans += num.Num
	}
}
