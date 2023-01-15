package main

import (
	"context"
	pb "example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.HelloServiceServer
}

func (*server) Hello(_ context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello " + request.GetName()

	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
