package client

import (
	"fmt"

	"strconv"

	"github.com/varshard/helloproto/addressbook"
	"github.com/varshard/helloproto/pingpong"
	"google.golang.org/grpc"
)

func StartAddressBookClient(args []string) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := addressbook.NewAddressBookClient(conn)

	if args[0] == "add" {
		p := addressbook.Person{
			FirstName: "First",
			LastName:  "Last",
		}

		resp, err := addressbook.AddPerson(client, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %v, message: %s\n", resp.GetSuccess(), resp.GetMessage())
	} else if args[0] == "list" {
		resp, err := addressbook.ListAddressBook(client)
		if err != nil {
			panic(err)
		}

		fmt.Printf("success: %d\n", len(resp.GetPersons()))
	}
}

func StartPingPongClient(args []string) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9001", opts...)
	if err != nil {
		panic(err)
	}

	client := pingpong.NewPingPongClient(conn)

	times, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	for i := 0; i < times; i++ {
		err = pingpong.StreamPing(client, int32(i))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Finish Pinging")
}
