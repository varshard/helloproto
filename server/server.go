package server

import (
	"fmt"
	"net"

	addressbook "github.com/varshard/helloproto/addressbook"

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
	fmt.Println("before serve")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("served")
}
