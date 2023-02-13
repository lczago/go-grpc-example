package controller

import (
	"context"
	"example/pb"
)

type LoginServer struct {
	pb.LoginServiceServer
}

func (*LoginServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if request.Name != "" {
		return &pb.LoginResponse{
			Authenticated: true,
		}, nil
	}
	return nil, nil
}
