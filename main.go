package main

import (
	"fmt"
	"os"
	"time"

	"tcp-connection/client"
	"tcp-connection/server"
)

var address = "localhost:6073"

// use main () to run client / server each in a separate process
func main() {
	errChan := make(chan error, 1)

	fmt.Println("main starting...")
	defer time.Sleep(time.Second)

	arg := os.Args[1]
	switch arg {
	case "server", "s":
		server.RunServer(address, errChan)
	case "client", "c":
		client.RunClient(address, errChan)
	default:
		fmt.Println("unknown argument. Usage: main.exe [client|server]")
		return
	}
	if err := <-errChan; err != nil {
		fmt.Println("Fatal error: ", err)
		return
	}

	fmt.Println("main done, all is OK :)")
}
