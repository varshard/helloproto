package client

import (
	"fmt"

	pb "github.com/varshard/helloproto/addressbook"
	api "github.com/varshard/helloproto/apis/addressbook"
	"google.golang.org/grpc"
)

func StartClient(args []string) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewAddressBookClient(conn)

	if args[0] == "add" {
		p := pb.Person{
			FirstName: "First",
			LastName:  "Last",
		}

		resp, err := api.AddPerson(client, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %v, message: %s\n", resp.GetSuccess(), resp.GetMessage())
	} else if args[0] == "list" {
		resp, err := api.ListAddressBook(client)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %d\n", len(resp.GetPersons()))
	}
}
