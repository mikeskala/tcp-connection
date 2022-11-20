package server

import (
	"bufio"
	"fmt"
	"net"
)

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
	msg, err := bufio.NewReader(conn).ReadString('.')
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("server got msg: ", msg)

	// test
	if msg != "hi there." {
		err = fmt.Errorf("server got unexpected message: %s", msg)
		errCh <- err
		return
	}

	// reply
	if _, err := fmt.Fprintf(conn, "ok."); err != nil {
		errCh <- err
		return
	}
	fmt.Println("sever replied")
	errCh <- nil
}
