package main

import (
	"log"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)



func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)

	var k int32 = 2

	N := in.Number


	for N > 1 {
		if N % k == 0{
			N = N / k

			stream.Send(&pb.PrimesResponse{
				Prime: k,
			})
			} else { 
				k++
			}

	}
	return nil
}

