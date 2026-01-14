package main

import (
	pb "github.com/israelmiranda/go-expert/grpc/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}
