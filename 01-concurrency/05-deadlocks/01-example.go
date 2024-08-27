package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go fn(wg)
	wg.Wait()
	fmt.Println("Done!")
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn invoked")
}
