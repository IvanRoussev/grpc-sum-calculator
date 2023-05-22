package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/IvanRoussev/grpc-sum-calculator/calculator/proto"
)

var addr string = "localhost:3000"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))


	if err != nil {
		log.Fatalf("Failed to Connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doCalculator(c)
	// doPrimes(c)
	doAvg(c)
}