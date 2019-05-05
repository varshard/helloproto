package main

import (
	"context"
	"time"
	"fmt"
	"os"

	pb "github.com/varshard/helloproto/unary/addressbook"
	"google.golang.org/grpc"
)

func StartAddressBookClient(args[] string) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	cmd := args[1]

	client := pb.NewAddressBookClient(conn)

	if cmd == "add" {
		p := pb.Person{
			FirstName: "First",
			LastName:  "Last",
		}

		resp, err := addPerson(client, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %v, message: %s\n", resp.GetSuccess(), resp.GetMessage())
	} else if cmd == "list" {
		resp, err := listAddressBook(client)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %d\n", len(resp.GetPersons()))
	}
}

func addPerson(client pb.AddressBookClient, person pb.Person) (*pb.AddPersonResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	// ads.Contacts = append(ads.Contacts, person)

	return client.AddPerson(ctx, &person)
}

func listAddressBook(client pb.AddressBookClient) (*pb.ListAddressBookResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	req := pb.ListAddressBookRequest{}
	return client.ListAddressBook(ctx, &req)
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func main() {
	StartAddressBookClient(os.Args)
}

