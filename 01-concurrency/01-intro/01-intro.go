package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of f1 through the scheduler
	f2()

	// poor man's synchronization (DO NOT DO THIS)
	time.Sleep(3 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
