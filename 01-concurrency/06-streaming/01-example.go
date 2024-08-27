package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
}

func genNos(ch chan<- int) {
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
}
