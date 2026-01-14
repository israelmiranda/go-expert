package main

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/israelmiranda/go-expert/grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("Primes func was invoked with: %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
