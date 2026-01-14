package main

import (
	"context"
	"io"
	"log"

	pb "github.com/israelmiranda/go-expert/grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "John",
	})
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", res.Result)
	}
}
