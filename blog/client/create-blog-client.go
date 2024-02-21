package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient, author, title, content string) string {
	log.Println("create blog was invoked , client side")

	data := &pb.Blog{
		AuthorId: author,
		Title:    title,
		Content:  content,
	}

	res, err := c.CreateBlog(context.Background(), data)

	if err != nil {
		log.Fatalf("Couldn't add blog, client side, error : %v", err)
	}

	log.Println("Log has been added")
	log.Printf("Blog Id added : %v\n", res.Id)

	return res.Id

}
