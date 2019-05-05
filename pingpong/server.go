package pingpong

import (
	"fmt"
	"io"
)

type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(stream PingPong_StartPingServer) error {
	fmt.Println("Ping Received")
	for {
		ping, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("Ping: %d\n", ping.Id)

		resp := Pong{
			Message: "Pong",
			Id:      ping.Id,
		}
		stream.Send(&resp)
	}

	return nil
}
