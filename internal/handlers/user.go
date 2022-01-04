package handlers

import (
	"context"
	pb "oblea/internal/rpc/proto"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

// create sayhello func
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloRequest, error) {
	return &pb.HelloRequest{Name: "Hello " + in.Name}, nil
}
