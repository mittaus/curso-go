package main

import (
	"context"
	"log"

	calcpb "example.com/grcp-calcpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := calcpb.NewCalculatorClient(cc)

	cReq := &calcpb.CalculatorRequest{
		X: 10,
		Y: 20,
	}

	cRes, err := c.Calculate(context.Background(), cReq)
	if err != nil {
		log.Printf("Server responded with error: %v\n", err)
	}

	log.Printf("Server responded with result: %v\n", cRes.GetResult())
}
