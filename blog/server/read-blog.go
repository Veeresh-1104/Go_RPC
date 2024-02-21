package main

import (
	"context"
	"log"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReadBlog(c context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Println("read blog server called")

	id := in.Id

	oid, ok := primitive.ObjectIDFromHex(id)

	if ok != nil {
		return nil, status.Errorf(codes.Internal, "Couldn't get the OID from ID passed")
	}

	data := &pb.Blog{}

	filter := bson.M{"_id": oid}

	res := Collection.FindOne(c, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "No Blog found")
	}

	return data, nil
}
