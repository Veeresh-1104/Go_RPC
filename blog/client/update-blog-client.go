package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
)

func UpdateBlog(ctx pb.BlogServiceClient) {

	_, err := ctx.UpdateBlog(context.Background(), &pb.Blog{
		AuthorId: "New Auth",
		Content:  "New Content",
		Title:    "New Title",
		Id:       "0x2dw342f32",
	})

	if err != nil {
		log.Fatalln("Could not get the update context")
	}

	log.Println("Blog Updated successfully ")

}
