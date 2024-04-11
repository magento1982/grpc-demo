package main

import (
	"context"

	pb "github.com/magento1982/grpc-demo/api/vs"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello the world",
	}, nil
}
