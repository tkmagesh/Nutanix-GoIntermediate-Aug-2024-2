package main

import (
	"fmt"
	"time"
)

func main() {
	dataCh := genNos()
	for no := range dataCh {
		fmt.Println(no)
	}
	fmt.Println("Done!")
}

func genNos() <-chan int {
	dataCh := make(chan int)
	timeoutCh := timeout(5 * time.Second)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-timeoutCh:
				break LOOP
			case dataCh <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(dataCh)
	}()
	return dataCh
}

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
