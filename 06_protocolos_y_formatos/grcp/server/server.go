package main

import (
	"context"
	"log"
	"net"

	calcpb "example.com/grcp-calcpb"
	"google.golang.org/grpc"
)

type calculatorServer struct{}

func (cs *calculatorServer) Calculate(ctx context.Context, cReq *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	x := cReq.GetX()
	y := cReq.GetY()

	log.Printf("Received request to calculate %d + %d\n", x, y)
	cRes := &calcpb.CalculatorResponse{
		Result: x + y,
	}

	return cRes, nil
}

func main() {
	addr := "localhost:50051"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	cServer := &calculatorServer{}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServer(s, cServer)

	log.Printf("Server listeing at %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
