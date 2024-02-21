package main

import (
	"context"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {

	oid, ok := primitive.ObjectIDFromHex(in.Id)

	if ok != nil {
		return nil, status.Errorf(codes.Internal, "Could not access the Object Id from the given ID")
	}

	_, err := Collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not Delete the Object")
	}

	return &emptypb.Empty{}, nil

}
