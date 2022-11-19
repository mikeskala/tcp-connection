package client

import (
	//"main"

	"fmt"
	"io/ioutil"
	"net"
	"sync"
)

func RunClient(address string, errCh chan error) {

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
	if _, err := fmt.Fprintf(conn, "hi there"); err != nil {
		errCh <- err
		return
	}
	fmt.Println("client sent msg")

	var replyWG sync.WaitGroup
	replyWG.Add(1)
	var success bool
	go func() {
		defer replyWG.Done()

		fmt.Println("client waiting for server reply")
		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			errCh <- err
			return
		}
		reply := string(buf[:])
		fmt.Println("client got reply: ", buf[:])
		fmt.Println("client got reply: ", reply)
		success = true
	}()
	replyWG.Wait()
	if !success {
		fmt.Println("client goroutine filed")
		return
	}

	fmt.Println("client exiting..")
	errCh <- nil
	defer fmt.Println("client exited")
}
