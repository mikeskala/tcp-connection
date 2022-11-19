package main

import (
	"fmt"
	"testing"

	"tcp-connection/client"
	"tcp-connection/server"
)

func TestConnBidirectional(t *testing.T) {

	errChan := make(chan error, 2)

	// run a client & server as threads
	go server.RunServer(address, errChan)
	go client.RunClient(address, errChan)

	fmt.Println("main test is waiting..")

	// select {
	// case err := <-errChan:
	// 	if err != nil {
	// 		t.Error("client error: ", err)
	// 	}
	// case err := <-chSrv:
	// 	if err != nil {
	// 		t.Error("server error: ", err)
	// 	}
	// case <-time.After(time.Second * 8):
	// 	t.Fatal("test stuck")
	// }

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			t.Error("client/server error: ", err)
		}
		fmt.Println("main test got message #", i)
	}

	fmt.Println("main test done")

}
