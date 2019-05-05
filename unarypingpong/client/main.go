package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/varshard/helloproto/unarypingpong/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendPing(client pb.PingPongClient) (*pb.Pong, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ping := pb.Ping{
		Id:      1,
		Message: "Ping",
	}
	pong, err := client.StartPing(ctx, &ping)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}
	fmt.Printf("Pong: %d, statusCode: %s\n", pong.Id, statusCode.String())
	return pong, err
}

func StartPingPongClient() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)

	_, err = SendPing(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish Pinging")
}

func main() {
	StartPingPongClient()
}
