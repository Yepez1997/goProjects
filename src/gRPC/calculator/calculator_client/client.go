package main


func  main() {
	fmt.Println("Calculator Client ...")

	cc, err := gprc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	// close the connection 
	defer cc.Close()

	c :=  calcalate.NewCalculateServiceClient(cc) 

	doUnaryCalculateSum(c)
}

func doUnaryCalculateSum(c calculatepb.CalculateServiceClient) {
	// request to send 
	req := calculatepb.CalculareRequest {
		Calculate: &calculate.Calculate {
			FirstInt: 10,
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