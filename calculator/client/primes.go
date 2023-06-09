package main

import (
	"context"
	"io"
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimesRequest{
		Number: 23478,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes(): %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("Primes: %v\n", msg.Prime)
	}
}