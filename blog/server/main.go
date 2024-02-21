package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

var Collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:pass@localhost:27017/?authSource=admin&authMechanism=SCRAM-SHA-1"))
	if err != nil {
		log.Fatalf("Error getting the client : %v\n", err)
	}

	Collection = client.Database("blogdb").Collection("blog")
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error in Listening to %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v\n", err)
	}

}
