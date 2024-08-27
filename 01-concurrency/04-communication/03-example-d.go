package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := add(100, 200)
	/*
		go func() {
			ch <- 10000
		}()
	*/
	result := <-ch
	fmt.Println("result :", result)
}

// producer
// returning a "receive-only" channel
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
