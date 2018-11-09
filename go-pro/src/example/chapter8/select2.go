// use select to simulate the timeout request
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9 * 11) 	// simulate the request (11s)
		timeout <- true		// channel receive the signal
	}()
	select {
	case <-timeout:
		fmt.Println("Completed!")
	case <- time.After(10 * time.Second):
		fmt.Println("Read time out!")	// 11s > 10s, so the request time out.
	}
}
