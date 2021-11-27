package main

import (
	"github.com/Garagator3000/learnGRPC/pkg/api"
	"github.com/Garagator3000/learnGRPC/pkg/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}


}
