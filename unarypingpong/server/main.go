package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/varshard/helloproto/unarypingpong/pingpong"
)

type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	fmt.Println("Ping Received")

	resp := pb.Pong{
		Id:      ping.Id,
		Message: ping.Message,
	}

	return &resp, nil
}

func StartPingPongServer() {
	server := PingPongServerImpl{}

	lis, err := net.Listen("tcp", "localhost:9000")

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongServer(grpcServer, &server)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("pingpong started")
}

func main() {
	StartPingPongServer()
}
