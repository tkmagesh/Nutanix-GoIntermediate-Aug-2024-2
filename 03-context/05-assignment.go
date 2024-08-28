/*
use context api for the below functionality
incorporate a timeout based cancellation as well (15 seconds)
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	dataCh := genNos(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...!")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range dataCh {
		fmt.Println(no)
	}
	fmt.Println("Done!")
}

func genNos(stopCh <-chan struct{}) <-chan int {
	dataCh := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			case dataCh <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(dataCh)
	}()
	return dataCh
}
