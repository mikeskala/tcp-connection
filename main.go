package main

import (
	"fmt"
	"os"

	"tcp-connection/client"
	"tcp-connection/server"
)

var address = "localhost:6073"

// use main () to run client / server each in a separate process
func main() {
	// THIS DOES NOT WORK WITH 2 SEPARATE PROCESSES
	// (I SEE THAT CLIENT & SERVER DON'T TALK ON THE SAME PORT)
	errChan := make(chan error, 1)

	arg := os.Args[1]
	switch arg {
	case "server", "s":
		server.RunServer(address, errChan)
	case "client", "c":
		client.RunClient(address, errChan)
	default:
		fmt.Println("unknown argument. Usage: main.exe [client|server]")
	}
	if err := <-errChan; err != nil {
		fmt.Println("Fatal error: ", err)
	}

	fmt.Println("main done.")
}
