package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
