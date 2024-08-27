package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(4)
	wg := &sync.WaitGroup{}
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Spinning up %d goroutines.. Hit ENTER to start!\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1)       // increment the wg counter by 1
		go fn(wg, id+1) // schedule the execution of f1 through the scheduler
	}
	wg.Wait() // block the execution until the counter becomes 0 (default value = 0)
	fmt.Println("Done!")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
