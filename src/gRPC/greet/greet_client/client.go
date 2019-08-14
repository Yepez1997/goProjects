package main

import (
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
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	// when whole main is done -> close the connection
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// place holder
	fmt.Printf("Created a client, %f", c)

}
