package pingpong

import (
	"fmt"
	"io"

	pb "github.com/varshard/helloproto/pingpong/pingpong"
)

type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(stream pb.PingPong_StartPingServer) error {
	fmt.Println("Start Ping")
	for {
		ping, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		resp := pb.Pong{
			Message: "Pong",
			Id:      ping.Id,
		}
		stream.Send(&resp)
	}

	return nil
}
