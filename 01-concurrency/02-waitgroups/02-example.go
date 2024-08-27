package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the wg counter by 1
	go f1(wg) // schedule the execution of f1 through the scheduler
	f2()

	wg.Wait() // block the execution until the counter becomes 0 (default value = 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")

}

func f2() {
	fmt.Println("f2 invoked")
}
