package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Yepez1997/goProjects/src/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

// all services - added on needed basis
type server struct{}

func main() {

	fmt.Print("Listening on port 50051 ...")
	// create connection; and port binding
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// the path to the protocol buffer
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
