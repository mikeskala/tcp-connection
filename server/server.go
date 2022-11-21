package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Msg struct {
	Body string
	Id   int
}

func RunServer(address string, errCh chan error) {
	defer fmt.Println("server exited")

	fmt.Println("server starting...")
	listener, err := net.Listen("tcp", address) // create server
	if err != nil {
		errCh <- err
		return
	}
	defer listener.Close()
	fmt.Println("server listening for incoming connections on:", address)

	conn, err := listener.Accept()
	if err != nil {
		errCh <- err
		return
	}
	defer conn.Close()
	fmt.Println("server connected to: Local:", conn.LocalAddr(), "Remote:", conn.RemoteAddr())

	// wait for message
	decoder := json.NewDecoder(bufio.NewReader(conn))
	var msg Msg
	err = decoder.Decode(&msg)
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("server got msg: ", msg)

	// test
	if msg.Body != "hi there" {
		err = fmt.Errorf("server got unexpected message body: %v", msg.Body)
		errCh <- err
		return
	}

	reply := `{"id": 2, "body": "ok"}`
	if _, err := fmt.Fprintf(conn, reply); err != nil {
		errCh <- err
		return
	}
	fmt.Println("sever replied")
	errCh <- nil
}
