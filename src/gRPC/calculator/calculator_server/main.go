package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	// "time"

	"github.com/Yepez1997/goProjects/src/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

// all services will bind to this struct
type server struct{}

// TOOD IMPLEMENT THIS
func (*server) CalculateSum(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	fmt.Printf("Calcualte Sum function was invoked with: %v", req)
	// look at pb def
	firstNumber := req.GetCalculate().GetFirstInt()
	secondNumber := req.GetCalculate().GetSecondInt()
	result := firstNumber + secondNumber
	// calculatepb response object
	res := &calculatorpb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func (*server) CalculatePrimes(req *calculatorpb.CalculateManyPrimesRequest, stream calculatorpb.CalculateService_CalculatePrimesServer) error {
	fmt.Printf("CalculatePrimes was invoked with: %v", req)
	/// get the number from the response
	number := req.GetNum().GetFirstNumber()
	k := int32(2)
	for number > 1 {
		if number%k == 0 {
			// build the object and send it
			res := &calculatorpb.CalculateManyPrimesResponse{
				Result: int32(k),
			}
			stream.Send(res)
			// time.Sleep(1000 * time.MilliSecond)
			number = int32(number / k)
		} else {
			k++
		}
	}
	return nil
}

func (*server) CalculateMax(stream calculatorpb.CalculateService_CalculateMaxServer) error {
	fmt.Printf("CalculateMax was invoked ...")
	currentMax := math.Inf(-1)
	// get the requests - onlt send the requets if the max is greater than the current max
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error occured receiving stream in Calculate Max: %v", err)
			return err
		}
		currentNumber := float64(request.GetNumber())
		// send if condtion persists
		if currentNumber > currentMax {
			stream.Send(&calculatorpb.CalculateMaxResponse{
				Result: int64(currentNumber),
			})
			// set new max
			currentMax = float64(currentNumber)
		}

	}
}

func (*server) CalculateAverage(stream calculatorpb.CalculateService_CalculateAverageServer) error {
	fmt.Println("Calculate Average Function was invoked")
	totalSum := int64(0)
	totalNum := int64(0)
	// go on forever until no more
	for {
		// receive stream
		request, err := stream.Recv()
		if err == io.EOF {
			// done and send back response - client should call receive and close
			return stream.SendAndClose(&calculatorpb.CalculateAverageResponse{
				Result: int64(totalSum / totalNum),
			})
		}
		if err != nil {
			log.Fatalf("Error occured in CalculateAverage Stream: %v", err)
		}
		number := request.GetNumber()
		totalSum += number
		totalNum++
	}
}

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
