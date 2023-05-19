package main

import (
	"context"
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)


func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	
	sum := in.Num1 + in.Num2

	return &pb.CalculatorResponse{
		// Result: in.first_num + in.second_num,
		Sum: sum,
	}, nil

}