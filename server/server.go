package server

import (
	"fmt"
	"net"

	"github.com/varshard/helloproto/addressbook"
	"github.com/varshard/helloproto/pingpong"

	"google.golang.org/grpc"
)

func StartAddressBookServer() {
	server := addressbook.AddressBookServerImpl{
		Contacts: make([]*addressbook.Person, 0),
	}

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	addressbook.RegisterAddressBookServer(grpcServer, &server)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("addressbook started")
}

func StartPingPongServer() {
	server := pingpong.PingPongServerImpl{}

	lis, err := net.Listen("tcp", "localhost:9001")

	grpcServer := grpc.NewServer()
	pingpong.RegisterPingPongServer(grpcServer, &server)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("pingpong started")
}
