package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "github.com/varshard/helloproto/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func StreamPing(client pb.PingPongClient, id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.StartPing(ctx)
	if err != nil {
		return err
	}

	ping := pb.Ping{
		Id:      id,
		Message: "Ping",
	}
	err = stream.Send(&ping)
	if err != nil {
		return err
	}

	pong, err := stream.Recv()

	statusCode := status.Code(err)
	fmt.Printf("Pong: %d, statusCode: %s\n", pong.Id, statusCode.String())
	return nil
}

func StartPingPongClient(times int32) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9001", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)

	for i := int32(0); i < times; i++ {
		err = StreamPing(client, i)
		if err != nil {
			panic(err)
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Finish Pinging")
}

func main() {
	args := os.Args
	times, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	StartPingPongClient(int32(times))
}
