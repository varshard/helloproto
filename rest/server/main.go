package main

import (
	"context"
	"fmt"
	pb "github.com/varshard/helloproto/rest/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

var contactId int32 = 1

func StartAddressBookServer() {
	fmt.Println("initialize GRPC Server")
	server := AddressBookServerImpl{
		Contacts: make(map[int32]pb.Person),
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
	Contacts map[int32]pb.Person
}

func (ads *AddressBookServerImpl) AddPerson(c context.Context, person *pb.Person) (*pb.AddPersonResponse, error) {
	ads.Contacts[contactId] = *person
	contactId++

	resp := pb.AddPersonResponse{
		Message: "contact added",
		Success: true,
	}

	return &resp, nil
}

func mapToArray(m map[int32]pb.Person) []*pb.Person {
	persons := make([]*pb.Person, 0)
	for k, _ := range m {
		person := m[k]
		persons = append(persons, &person)
	}

	return persons
}

func (ads *AddressBookServerImpl) ListAddressBook(c context.Context, in *pb.ListAddressBookRequest) (*pb.ListAddressBookResponse, error) {

	return &pb.ListAddressBookResponse{
		Success: true,
		Persons: mapToArray(ads.Contacts),
	}, nil
}

func (ads *AddressBookServerImpl) GetPerson(c context.Context, in *pb.GetPersonRequest) (*pb.Person, error) {
	person, exist := ads.Contacts[in.Id]

	if !exist {
		status.Error(codes.NotFound, "Contact not found")
	}

	return &person, nil
}

func main() {
	StartAddressBookServer()
}
