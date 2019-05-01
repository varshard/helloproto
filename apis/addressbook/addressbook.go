package addressbook

import (
	"context"
	"time"

	pb "github.com/varshard/helloproto/addressbook"
)

func AddPerson(client pb.AddressBookClient, person pb.Person) (*pb.AddPersonResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	return client.AddPerson(ctx, &person)
}

func ListAddressBook(client pb.AddressBookClient) (*pb.ListAddressBookResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	req := pb.ListAddressBookRequest{}
	return client.ListAddressBook(ctx, &req)
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
