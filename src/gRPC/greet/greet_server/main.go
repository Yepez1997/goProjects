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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// all services - added on needed basis
type server struct{}
type server2 struct{}

// unary
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

func (*server) GreetGoodbye(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GoodbyeResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	// in other words get te buffer bytes
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	// create new protobuffer request
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	req = &greetpb.GoodbyeRequest{
		Result: firstName,
	}

	result, err = (*server2).Goodbye(context.Background(), req)
	return result, err
}

func (*server2) Goodbye(ctx context.Context, req *greetpb.GoodbyeRequest) (*greetpb.GoodbyeResponse, error) {
	result := "Goodbye " + firstName
	res := &greetpb.GoodbyeResponse{
		Result: result,
	}
	return res

}

// server streaming
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

// client streaming
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked... ")
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

// bidirectional streaming
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	// receive the message
	// send the message back
	// remember to check for errors
	fmt.Println("Greet Everyone function was invoked ...")
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// no more streaming
			return nil
		}
		if err != nil {
			return err
		}
		// get the request
		firstName := request.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		// check for error and send back the response from that particular request
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error sending the data back from GreetEveryone Server: %v", err)
			return err
		}

	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Println("Greet function was invoked ...")
	// sleep the function for 3 seconds and want to check if the error has been cancelled
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client cancelled the request
			fmt.Println("The client cancled the request")
			// throw the deadline cancelled error
			return nil, status.Error(codes.DeadlineExceeded, "The client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	// get the first name from the request
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadlineResponse{
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
	// adding authentication
	// server certificate
	certFile := "../../ssl/server.crt"
	// .pem file is a file grpc can read
	keyFile := "../../ssl/server.pem"
	cred, SSLerr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if SSLerr != nil {
		fmt.Printf("Error loading credentials: %v", err)
		return
	}
	// pass in the credentials to the grpc call
	s := grpc.NewServer(grpc.Creds(cred))
	// the path to the protocol buffer
	//greetpb.RegisterGreetServiceServer(s, &server{})
	greetpb.RegisterGoodbyeServiceServer(s, &server2{})
	// adding reflection service to the server
	//reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
