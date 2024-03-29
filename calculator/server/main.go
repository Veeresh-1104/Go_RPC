package main

import (
	"log"
	"net"

	pb "github.com/Veeresh-1104/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.AddTwoNumbrsServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error in Listening to %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterAddTwoNumbrsServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v\n", err)
	}

}
