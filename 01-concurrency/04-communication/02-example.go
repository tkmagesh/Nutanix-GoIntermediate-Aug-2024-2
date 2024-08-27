package main

import (
	"fmt"
	"sync"
)

func main() {
	// result := add(100, 200)
	// communicate by sharing memory
	var result int
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = add(100, 200)
	}()
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int) int {
	return x + y
}
