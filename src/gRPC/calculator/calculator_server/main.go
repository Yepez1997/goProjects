package main


import (
	"context"
	"fmt"
	"log"
	"net"
)

// all services will bind to this struct 
type server struct{}

func main() {
	fmt.Println("Listening to port 50051 ...")
	// below is boiler plate code 
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// from .proto.go file
	calculatorpb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}