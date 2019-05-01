package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/varshard/helloproto/addressbook"
	"google.golang.org/grpc"
)

type AddressBookServer struct {
	Contacts []*pb.Person
}

func (ads *AddressBookServer) AddPerson(c context.Context, person *pb.Person) (*pb.AddPersonResponse, error) {
	ads.Contacts = append(ads.Contacts, person)

	return &pb.AddPersonResponse{
		Message: "contact added",
		Success: true,
	}, nil
}

func (ads *AddressBookServer) ListAddressBook(c context.Context, in *pb.ListAddressBookRequest) (*pb.ListAddressBookResponse, error) {
	return &pb.ListAddressBookResponse{
		Success: true,
		Persons: ads.Contacts,
	}, nil
}

func StartServer() {
	server := AddressBookServer{
		Contacts: make([]*pb.Person, 0),
	}

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAddressBookServer(grpcServer, &server)
	fmt.Println("before serve")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("served")
}
