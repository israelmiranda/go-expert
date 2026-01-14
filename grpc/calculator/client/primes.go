package main

import (
	"context"
	"io"
	"log"

	pb "github.com/israelmiranda/go-expert/grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes was invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Println(res.Result)
	}
}
