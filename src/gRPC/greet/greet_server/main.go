package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Yepez1997/goProjects/src/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

// all services - added on needed basis
type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	// create a message that the function was invoked
	fmt.Printf("Greet function was invoked with %v", req)
	// in other words get te buffer bytes 
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	// create new protobuffer request
	res := &greetpb.GreetingResponse{
		Result: result,
	}
	return res, nil
}

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
