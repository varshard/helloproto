package main

import (
	"fmt"

	"github.com/varshard/helloproto/client"
	"github.com/varshard/helloproto/server"

	"os"
)

func main() {
	args := os.Args
	fmt.Printf("args %s\n", args[1])
	switch args[1] {
	case "server":
		{
			server.StartServer()
		}
	case "client":
		{
			client.StartClient(args[2:])
		}
	}
}
