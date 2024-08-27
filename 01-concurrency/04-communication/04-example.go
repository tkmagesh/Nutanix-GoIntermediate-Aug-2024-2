/* channel behavior */
package main

import (
	"fmt"
)

/*
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
	wg.Wait()
}
*/

func main() {
	ch := make(chan int)
	doneCh := make(chan struct{})
	go func() {
		data := <-ch
		fmt.Println(data)
		doneCh <- struct{}{}
	}()
	ch <- 100
	<-doneCh // blocking until the goroutine completes
}
