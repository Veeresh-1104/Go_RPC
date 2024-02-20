package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
)

func MaxApiDualStream(c pb.AddTwoNumbrsClient) {
	log.Println("MaxAPI client side invoked")

	stream, err := c.MaxApiBiStream(context.Background())

	reqs := []*pb.MaxApiRequest{
		{Num: -1},
		{Num: -10},
		{Num: 12},
		{Num: 3400},
		{Num: 4399},
		{Num: 2000},
		{Num: 69},
		{Num: 19},
		{Num: 76},
	}
	if err != nil {
		log.Println("Couldn't initialise a stream")
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sent value : %v", req.Num)
			stream.Send(req)
			time.Sleep(2 * time.Second)
		}

		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving from the server")
				break
			}

			log.Printf("Max Value : %v\n", res.Max)
		}

		close(waitc)
	}()

	<-waitc

}
