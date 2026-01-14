package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/israelmiranda/go-expert/grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet func was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.FirstName),
	}, nil
}
