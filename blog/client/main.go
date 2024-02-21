package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Erroring creating the client connection %v\n", err)
	}

	client := pb.NewBlogServiceClient(conn)

	CreateBlog(client, "Veeresh", "Blog - 1", "Hello this is my first blog")
	// CreateBlog(client, "Hola", "Blog - 2", "Hello this is my first blog")
	// CreateBlog(client, "bhola", "Blog - 3", "Hello this is my first blog")

	defer conn.Close()

}
