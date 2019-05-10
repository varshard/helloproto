package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	pb "github.com/varshard/helloproto/bidirectionalstream/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func StreamPing(client pb.PingPongClient, times int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	startPingClient, err := client.StartPing(ctx)
	if err != nil {
		return err
	}

	// For bidirectional streaming, request and response don't necessary come in sync.
	// So, I use goroutine for both streaming and listening.

	// Stream Ping to the server.
	// Use goroutine here. So, client can send a new message asynchronously.
	go func() {
		for i := int32(0); i < times; i++ {
			ping := pb.Ping{
				Id:      i,
				Message: "Ping",
			}
			fmt.Printf("Ping: %d\n", i)
			err := startPingClient.Send(&ping)
			if err != nil {
				fmt.Printf("Ping error: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}

		// Tell the server it has done sending messages to the server.
		// The client can send many messages before calling CloseSend().
		fmt.Println("close send")
		startPingClient.CloseSend()
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	// Listen to Pong stream from the server.
	// Use goroutine here. So, client can receive a new message asynchronously.
	go func() {
		defer wg.Done()
		for {
			pong, err := startPingClient.Recv()

			// Done listening once client received EOF
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else if err != nil {
				fmt.Printf("Error pong: %v\n", err)
				break
			}

			statusCode := status.Code(err)
			fmt.Printf("Pong: %d, %s, statusCode: %s\n", pong.Id, pong.Message, statusCode.String())

			// Force receive to go out of sync with the request to demonstrate
			// that request and response are independent of each other.
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()

	return nil
}

func StartPingPongClient(times int32) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9001", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)

	err = StreamPing(client, times)
	if err != nil {
		panic(err)
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
