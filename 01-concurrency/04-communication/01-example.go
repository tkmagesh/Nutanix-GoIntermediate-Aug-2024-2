package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
/*
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println("result :", result)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
*/

func main() {
	var result int
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, &result) //communicate by sharing memory
	wg.Wait()
	fmt.Println("result :", result)
}

func add(x, y int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	*result = x + y
}
