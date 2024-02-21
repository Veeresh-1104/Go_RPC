package main

import (
	"context"
	"fmt"

	pb "github.com/Veeresh-1104/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(empt *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {

	curr, err := Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown Error : %v", err))
	}

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &pb.Blog{}

		err := curr.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Some Unknown error : %v", err))
		}

		stream.Send(data)

	}

	return nil

}
