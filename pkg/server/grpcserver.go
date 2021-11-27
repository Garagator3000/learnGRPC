package server

import (
	"context"
	"github.com/Garagator3000/learnGRPC/pkg/api"
)

type GRPCServer struct {}

func (server *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
