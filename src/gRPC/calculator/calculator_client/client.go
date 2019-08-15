package main

import (
	"context"
	"fmt"
	"log"

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

	doUnaryCalculateSum(c)
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
