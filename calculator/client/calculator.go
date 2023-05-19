package main

import (
	"context"
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)


func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCaclulator was invoked")
	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		Num1: 2,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Could not Calculate: %v\n", err)
	} 

	log.Printf("Calculated Sum is: %v\n", res.Sum)
}