package api

import (
	"context"
)

//protoc --go_out=. --go-grpc_out=. adder.proto

type GRPCServer struct{}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	panic("implement me")
}

func (s *GRPCServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return &AddResponse{Result: req.GetX() + req.GetY()}, nil
}
