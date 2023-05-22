package main

import (
	"context"
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)


func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg has been invoked")

	reqs := []*pb.AvgRequest {
		{ Avg: 3 }, 
		{ Avg: 5 }, 
		{ Avg: 9 }, 
		{ Avg: 54 }, 
		{ Avg: 23}, 
	}

	var length = len(reqs)

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAvg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("AVG: %v\n", float32(res.Avg) / float32(length))


}