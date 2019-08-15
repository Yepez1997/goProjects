package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
	// calculatepb response
	res := &calculatorpb.CalculateResponse{
		Result: result,
	}
	return res, nil
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
