package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"

	"tcp-connection/client"
)

var address = "localhost:6073"
var errServer error

func server(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("server starting...")
	listener, err := net.Listen("tcp", address) // create server
	if err != nil {
		errServer = err
		return
	}
	defer listener.Close()
	fmt.Println("server listening to:", address)

	conn, err := listener.Accept()
	if err != nil {
		errServer = err
		return
	}
	defer conn.Close()
	fmt.Println("server connected to: Local:", conn.LocalAddr(), "Remote:", conn.RemoteAddr())

	// wait for message
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		errServer = err
		return
	}
	msg := string(buf[:])
	fmt.Println("server got msg: ", msg)

	// test
	if msg != "hi there" {
		err = fmt.Errorf("server got unexpected message: %s", msg)
		errServer = err
		return
	}

	// reply
	if _, err := fmt.Fprintf(conn, "ok"); err != nil {
		errServer = err
		return
	}
	fmt.Println("sever replied")
}

func main() {
	// THIS DOES NOT WORK WITH 2 SEPARATE PROCESSES
	// (I SEE THAT CLIENT & SERVER DON'T TALK ON THE SAME PORT)
	var wg sync.WaitGroup // not in use here
	chCli := make(chan error, 1)

	arg := os.Args[1]
	switch arg {
	case "server", "s":
		server(&wg)
	case "client", "c":
		client.RunClient(address, chCli)
		if err := <-chCli; err != nil {
			fmt.Println("Fatal error: ", err)
		}
	default:
		fmt.Println("unknown argument. Usage: main.exe [client|server]")
	}

	fmt.Println("main done.")
}
