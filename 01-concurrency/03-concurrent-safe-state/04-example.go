package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count atomic.Int64

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go Increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", count.Load())
}

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count.Add(1)

}
