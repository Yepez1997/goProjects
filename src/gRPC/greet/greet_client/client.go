package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Yepez1997/goProjects/src/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World from Client")
	// create a connection to the server
	// by def grpc has ssl
	// once going in to production - remove
	cc, err := grpc.Dial("localxwhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	// when whole main is done -> close the connection
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnaryGreet(c)
	doServerStreaming(c)
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


// do Server streaming 
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server stream RPC ...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: greetpb.Greeting{
			FirstName: "Aureliano",
			LastName: "Yepez",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not process response: %v", res)
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