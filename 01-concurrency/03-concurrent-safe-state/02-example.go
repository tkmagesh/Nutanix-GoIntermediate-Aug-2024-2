package main

import (
	"fmt"
	"sync"
)

// custom type for concurrent safe state manipulation
type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) Add(delta int) {
	c.mutex.Lock()
	{
		c.count += delta
	}
	c.mutex.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go Increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", counter.count)
}

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)

}
