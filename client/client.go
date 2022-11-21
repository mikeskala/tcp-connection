package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"tcp-connection/server"
)

func RunClient(address string, errCh chan error) {
	defer fmt.Println("client exited")

	fmt.Println("client starting...")
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		errCh <- err
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) // connect to server (address)
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("client connected to: Local:", conn.LocalAddr(), "Remote:", conn.RemoteAddr())
	defer conn.Close()

	// write a message:
	payload, _ := json.Marshal(server.Msg{Body: "hi there", Id: 1})
	_, err = conn.Write(payload)
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("client sent msg, waiting for server reply")

	// read a json message (from TCP)
	decoder := json.NewDecoder(bufio.NewReader(conn))
	var reply server.Msg
	err = decoder.Decode(&reply)
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("client got reply: ", reply)

	errCh <- nil
}
