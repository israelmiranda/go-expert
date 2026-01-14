package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/israelmiranda/go-expert/grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes was invoked with: %v\n", in)

	for i := range 10 {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
