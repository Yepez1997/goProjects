package main

import (
	"context"
	"fmt"
	"log"
	"io"

	"github.com/Yepez1997/goProjects/src/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client ...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	// close the connection
	defer cc.Close()

	c := calculatorpb.NewCalculateServiceClient(cc)
	doServerPrimeStreaming(c)
	//doUnaryCalculateSum(c)
}

func doUnaryCalculateSum(c calculatorpb.CalculateServiceClient) {
	// request to send
	req := &calculatorpb.CalculateRequest{
		Calculate: &calculatorpb.Calculate{
			FirstInt:  10,
			SecondInt: 3,
		},
	}

	// send the calculate sum
	res, err := c.CalculateSum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// print the result
	log.Printf("Response from CalculateSum: %v", res.Result)

}



func doServerPrimeStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a server stream RPC ...")
	// format the request 
	req  := &calculatorpb.CalculateManyPrimesRequest{
		Num: &calculatorpb.Number{
			FirstNumber: 120,
		},
	}

	resStream, err := c.CalculatePrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not process response: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		// if at the end 
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading the stream %v", err)

		}
		// print response
		log.Printf("Response from Greet many times %v", msg.GetResult())
	}



}
