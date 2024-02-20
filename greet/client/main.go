package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Veeresh-1104/grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Erroring creating the client connection %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	// DoGreet(c)
	// GreetManyTimes(c)
	DoLongGreet(c)

	defer conn.Close()

}
