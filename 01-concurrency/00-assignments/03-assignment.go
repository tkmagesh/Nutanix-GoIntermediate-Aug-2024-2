/* modify the below so that the data is continuously produced until the user hits ENTER key */
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
	timeoutCh := time.After(5 * time.Second)
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
