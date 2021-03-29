package main

import (
	"google.golang.org/grpc"
	"grpc/api"
	"log"
	"net"
)

func main() {

	s := grpc.NewServer()
	srv := &api.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8185")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
