package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	// get the name
	// iterate n times
	// create the response object
	// send the object to the stream
	// optional -> timeout
	fmt.Println("GreetManyTimes function was invoked: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 5; i++ {
		msg := "Hello " + firstName + " " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: msg,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)

	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked")
	// result we are going to send back
	result := ""
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// END OF THE FILE
			// NEED TO SEND BACK THE RESPONSE
			return stream.SendAndClose(&greetpb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error occured during the steam: %v", err)
		}
		// deserialize the request
		firstName := msg.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "

	}
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
