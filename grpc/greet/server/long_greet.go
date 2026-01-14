package main

import (
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/israelmiranda/go-expert/grpc/greet/proto"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Printf("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
