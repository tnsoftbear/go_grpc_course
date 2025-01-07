package main

import (
	"github.com/cshep4/grpc-course/module5-exercise/internal/token"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const bindAddr = ":50051"

func main() {
	grpcServer := grpc.NewServer()
	tokenService := token.Service{}
	proto.RegisterTokenServiceServer(grpcServer, tokenService)
	lis, err := net.Listen("tcp", bindAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("Starting gRPC server on address: %s", bindAddr)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
