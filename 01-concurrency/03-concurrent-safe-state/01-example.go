package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go Increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()

}
