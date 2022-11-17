package main

import (
	"fmt"
	"sync"
	"testing"

	"tcp-connection/client"
)

func TestConnBidirectional(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)
	chCli := make(chan error, 1)
	// chSrv := make(chan error)

	// run a client & server threads
	go server(&wg)
	go client.RunClient(address, chCli)

	fmt.Println("main is waiting..")
	wg.Wait()

	// select {
	// case err1 := <-chCli:
	// 	t.Fatal("client error: ", err1)
	// case err2 := <-chSrv:
	// 	t.Fatal("server error: ", err2)
	// }

	if err := <-chCli; err != nil {
		t.Error("client error: ", err)
	}
	if errServer != nil {
		t.Error("server error: ", errServer)
	}
	fmt.Println("done.")

}
