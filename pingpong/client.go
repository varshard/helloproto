package pingpong

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/status"
)

func StreamPing(client PingPongClient, id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.StartPing(ctx)
	if err != nil {
		return err
	}

	ping := Ping{
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
