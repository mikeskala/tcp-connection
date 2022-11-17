package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"testing"
)

var address = "localhost:6073"
var message = "Hi there!"
var errClient, errServer error

func client(wg *sync.WaitGroup) {
	defer wg.Done()

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		errClient = err
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) // connect to server (address)
	if err != nil {
		errClient = err
		return
	}
	defer conn.Close()
	fmt.Println("client connected")

	// write a message: c.NetConn.Write(msg)
	if _, err := fmt.Fprintf(conn, message); err != nil {
		errClient = err
		return
	}
	fmt.Println("client sent msg")
}

func server(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", address) // create server
	if err != nil {
		errServer = err
		return
	}
	fmt.Println("server listening")
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		errServer = err
		return
	}
	defer conn.Close()
	fmt.Println("server accepting")

	// wait for message
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		errServer = err
		return
	}

	msg := string(buf[:])
	fmt.Println("server got msg: ", msg)

	if msg != (message) {
		err = fmt.Errorf("server got unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		errServer = err
		return
	}
}

func TestConnBD(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	go client(&wg)

	go server(&wg)

	fmt.Println("main is waiting..")
	wg.Wait()

	if errClient != nil {
		t.Fatal(errClient)
	}
	if errServer != nil {
		t.Fatal(errServer)
	}
	fmt.Println("main done.")

}
