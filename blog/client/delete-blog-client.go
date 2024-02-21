package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
)

func DeleteBlogById(c pb.BlogServiceClient) {

	id := "0xdmkwlefweinf"

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("the record could not be deleted")
	}

	log.Println("The record is deleted successfully deleted")

}
