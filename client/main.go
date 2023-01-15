package main

import (
	"context"
	"example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "TRANSDATA",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("Erro during the execution: %v", err)
	}

	log.Println(res.GetMsg())
}
