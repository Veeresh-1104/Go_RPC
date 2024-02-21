package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	log.Println("Update Blog is called")
	id := in.Id

	oid, ok := primitive.ObjectIDFromHex(id)

	if ok != nil {
		return nil, status.Errorf(codes.Internal, "Couldn't get the OID from ID passed")
	}

	filter := bson.M{"_id": oid}
	data := &pb.Blog{
		AuthorId: "hola",
		Title:    "Title",
		Content:  in.Content,
	}

	res, err := Collection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		log.Fatalf("Error while updating the records \n")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "No Document found with the given ID")
	}

	return &emptypb.Empty{}, nil

}
