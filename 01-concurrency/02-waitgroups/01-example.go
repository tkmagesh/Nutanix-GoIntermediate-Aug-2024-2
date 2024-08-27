package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the wg counter by 1
	go f1()   // schedule the execution of f1 through the scheduler
	f2()

	wg.Wait() // block the execution until the counter becomes 0 (default value = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
