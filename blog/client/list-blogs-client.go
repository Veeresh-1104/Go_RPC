package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient) {

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalln("Error while initialising the stream")
	}

	for {
		doc, err := stream.Recv()

		if err != io.EOF {
			log.Println("Reached the end of the Stream data")
			break
		}

		if err != nil && err != io.EOF {
			log.Fatalf("Unexpected Error : %v\n", err)
			break
		}

		log.Printf("the document recevived : %v\n", doc)
	}

}
