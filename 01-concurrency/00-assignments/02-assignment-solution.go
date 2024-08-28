/*
rewrite to follow "share memory by communicating"
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primesCh := findPrimes(1000, 2000)
	for primeNo := range primesCh {
		fmt.Println("Prime :", primeNo)
	}
	fmt.Println("Done")
}

func findPrimes(start, end int) <-chan int {
	// share memory by communicating
	var primesCh chan int = make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if isPrime(no) {
					primesCh <- no
				}
			}()
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
