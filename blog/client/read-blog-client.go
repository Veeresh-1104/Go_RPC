package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
)

func ReadBlog(c pb.BlogServiceClient) {
	log.Println("blog Read was invoked")

	req := &pb.BlogId{Id: "0xfewfefw23fe"}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while Reading blog client side : %v\n", err)
	}

	log.Printf("The blog : %v", res)
}
