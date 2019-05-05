package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/varshard/helloproto/unary/addressbook"
)

func StartAddressBookServer() {
	fmt.Println("initialize GRPC Server")
	server := AddressBookServerImpl{
		Contacts: make([]*pb.Person, 0),
	}

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterAddressBookServer(grpcServer, &server)

	if err = grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("addressbook started")
}

type AddressBookServerImpl struct {
	Contacts []*pb.Person
}

func (ads *AddressBookServerImpl) AddPerson(c context.Context, person *pb.Person) (*pb.AddPersonResponse, error) {
	ads.Contacts = append(ads.Contacts, person)

	resp := pb.AddPersonResponse{
		Message: "contact added",
		Success: true,
	}

	return &resp, nil
}

func (ads *AddressBookServerImpl) ListAddressBook(c context.Context, in *pb.ListAddressBookRequest) (*pb.ListAddressBookResponse, error) {
	return &pb.ListAddressBookResponse{
		Success: true,
		Persons: ads.Contacts,
	}, nil
}

func main() {
	StartAddressBookServer()
}
