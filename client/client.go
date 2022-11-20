package client

import (
	"bufio"
	"fmt"
	"net"
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

	// write a message: conn.Write(msg)
	if _, err := fmt.Fprintf(conn, "hi there."); err != nil {
		errCh <- err
		return
	}
	fmt.Println("client sent msg, waiting for server reply")

	reply, err := bufio.NewReader(conn).ReadString('.')
	// buf, err := ioutil.ReadAll(conn)
	if err != nil {
		errCh <- err
		return
	}
	fmt.Println("client got reply: ", reply)

	errCh <- nil
}
