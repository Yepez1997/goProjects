package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0.50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()

}
