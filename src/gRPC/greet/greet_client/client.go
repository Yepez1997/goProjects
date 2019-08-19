package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Yepez1997/goProjects/src/gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello World from Client")
	// create a connection to the server
	// by def grpc has ssl
	// once going in to production - remove
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	// when whole main is done -> close the connection
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnaryGreet(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBidirectionalStreaming(c)
	doUnaryWithDeadline(c)

}

// doUnaryGreet - unary request -> unary response for greet service
func doUnaryGreet(c greetpb.GreetServiceClient) {
	// ideally should pass in first name and last name as a variable
	// set up the request -> the object
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Aureliano",
			LastName:  "Yepez",
		},
	}

	// call the greet function on the server
	res, err := c.Greet(context.Background(), req)
	// check for err
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// print the result
	log.Printf("Response from Greet: %v", res.Result)

}

// doClientStreaming - send many client streams (by default with http 2)
func doClientStreaming(c greetpb.GreetServiceClient) {
	// define the requests to send
	// grab the stream and send all the requests to that stream
	fmt.Println("Starting to do a client stream RPC ...")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aureliano",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Adrian",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Alexander",
			},
		},
	}
	// set up the requests
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error occured in LongGreet: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending request: %v", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	// close the stream once done sending and receive the response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error occured while closing the stream %v", err)
	}
	fmt.Printf("Long Greet response: %v", res)
}

// do Server streaming
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server stream RPC ...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Aureliano",
			LastName:  "Yepez",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not process response: %v", err)
	}
	// print the result
	// keep looping until you reach the end
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream %v", err)
		}
		log.Printf("Response from Greet many times %v", msg.GetResult())
	}

}

// do BidirectionalStreaming - makes use of go routines
func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bidirectional Stream RPC  ...\n")

	// create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error occured receiving the stream in GreetEveryone: %v", err)
	}

	// the request we want to send
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aureliano",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Adrian",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Alexander",
			},
		},
	}

	waitc := make(chan struct{})
	// send a bunch of messages to the client in their own go routines
	go func() {
		// send every request
		for _, req := range requests {
			fmt.Printf("Sending Request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Millisecond)
		}
		// close the stream once done
		stream.CloseSend()
	}()
	// receive a bunch of mesages from the client
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error in receiving stream: %v", err)
				break
			}
			fmt.Printf("Response: %v\n", req.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
// doUnaryWithDeadline - unary api call with deadline requests 
func doUnaryWithDeadline(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC call with Deadline")
	// define the rpc request 
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Aureliano",
			LastName: "Yepez",
		}
	}
	// ctx := context.Background()
	// pass in context with time out by default
	ctc, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	// call the function 
	res, err := greetpb.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromErr(err)
		// recall if ok -> was grpc error 
		// if the context is hit -> should handle in the error 
		if ok {
			// print out of error codes 
			fmt.Println("Error message: %v", statusErr.Message())
			fmt.Println("Error code: %v", statusErr.Code())
			if statusErr.Code() ==  codes.DeadlineExceeded {
				fmt.Println("Time out was hit; deadline was exceeded")
			} else {
				fmt.Println("Unexpected error: %v", statusError)
			}
		} else {
				fmt.Println("Error in doUnaryWithDeadline, %v", err)
		}
	}
	// print the results 
	fmt.Println("Response UnaryWithDeadline: %v", res.GetResult())
}
