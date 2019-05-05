package main

import (
	"github.com/varshard/helloproto/client"
	"github.com/varshard/helloproto/server"

	"os"
)

func main() {
	args := os.Args
	switch args[1] {
	case "server":
		{
			switch args[2] {
			case "addressbook":
				server.StartAddressBookServer()
			case "pingpong":
				server.StartPingPongServer()
			}
		}
	case "client":
		{
			switch args[2] {
			case "addressbook":
				client.StartAddressBookClient(args[3:])
			case "pingpong":
				client.StartPingPongClient(args[3:])
			}
		}
	}
}
