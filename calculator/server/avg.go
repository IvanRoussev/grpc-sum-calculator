package main

import (
	"io"
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)


func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg Function has been invoked")

	var res int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: res,
			})
		}


		if err != nil {
			log.Fatalf("Error while reading client Stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += req.Avg
	}
}