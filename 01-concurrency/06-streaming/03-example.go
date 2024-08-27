package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done")

}

func genNos(ch chan<- int) {
	var count int
	count = rand.Intn(20)
	fmt.Println("[genNos] count :", count)
	for no := range count {
		ch <- (no + 1) * 10
	}
	close(ch)
}
