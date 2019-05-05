package main

import (
	"fmt"
	"io"
	"net"
	"google.golang.org/grpc"

	pb "github.com/varshard/helloproto/pingpong"
)

type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(stream pb.PingPong_StartPingServer) error {
	fmt.Println("Ping Received")
	for {
		ping, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("Ping: %d\n", ping.Id)

		resp := pb.Pong{
			Message: "Pong",
			Id:      ping.Id,
		}
		stream.Send(&resp)
	}

	return nil
}

func StartPingPongServer() {
	server := PingPongServerImpl{}

	lis, err := net.Listen("tcp", "localhost:9001")

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
