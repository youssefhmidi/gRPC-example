package main

import (
	"log"
	"net"

	"github.com/youssefhmidi/gRPC-example/services"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting gRPC server")

	gRPCServer := grpc.NewServer()
	l, err := net.Listen("tcp", ":54004")
	if err != nil {
		log.Fatalf("Couldn't start tcp server: %s", err)
		return
	}

	server := &Server{}
	services.RegisterMsgServerServer(gRPCServer, server)

	log.Println("Starting server at port 54004")
	gRPCServer.Serve(l)
}
