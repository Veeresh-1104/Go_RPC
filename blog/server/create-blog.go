package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create Blog was invoked : %v", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Content:  in.Content,
		Title:    in.Title,
	}

	res, err := Collection.InsertOne(ctx, data)

	if err != nil {
		log.Fatalf("Could not insert blog into database %v\n", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not insert blog into database %v\n", err))
	}

	Oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "could not convert to OID\n")
	}

	//.hex() converts to string
	return &pb.BlogId{
		Id: Oid.Hex(),
	}, nil

}
