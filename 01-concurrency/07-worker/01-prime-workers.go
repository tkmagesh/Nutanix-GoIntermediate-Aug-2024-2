/*
rewrite to follow "share memory by communicating"
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primesCh := findPrimes(1000, 2000, 20)
	for primeNo := range primesCh {
		fmt.Println("Prime :", primeNo)
	}
	fmt.Println("Done")
}

func findPrimes(start, end int, workerCount int) <-chan int {
	// share memory by communicating
	var primesCh chan int = make(chan int)
	dataCh := dataProducer(start, end)
	go func() {
		wg := &sync.WaitGroup{}
		for id := range workerCount {
			wg.Add(1)
			go primeWorker(id+1, wg, dataCh, primesCh)
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func dataProducer(start, end int) <-chan int {
	var dataCh chan int = make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func primeWorker(id int, wg *sync.WaitGroup, dataCh <-chan int, primeCh chan<- int) {
	fmt.Printf("worker [%d] started..!\n", id)
	defer wg.Done()
	for no := range dataCh {
		if isPrime(no) {
			primeCh <- no
		}
	}
	fmt.Printf("worker [%d] completed..!\n", id)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
