package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//doServerPrimeStreaming(c)
	//doUnaryCalculateSum(c)
	doClientStreaming(c)
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
	req := &calculatorpb.CalculateManyPrimesRequest{
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

// clientStreaming API for average api
func doClientStreaming(c calculatorpb.CalculateServiceClient) {
	// set up the requests
	fmt.Println("Client Streaming API ... ")
	// should spit out 4 from the server
	requests := []*calculatorpb.CalculateAverageRequest{
		&calculatorpb.CalculateAverageRequest{
			Number: 4,
		},
		&calculatorpb.CalculateAverageRequest{
			Number: 5,
		},
		&calculatorpb.CalculateAverageRequest{
			Number: 3,
		},
		&calculatorpb.CalculateAverageRequest{
			Number: 3,
		},
		&calculatorpb.CalculateAverageRequest{
			Number: 5,
		},
	}

	// get the stream
	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("Error occuered in Calculate Average: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	// close the stream
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error occured in response for Calculate Averege: %v", err)
	}
	fmt.Printf("CalculateAverage Response: %v", response)

}
