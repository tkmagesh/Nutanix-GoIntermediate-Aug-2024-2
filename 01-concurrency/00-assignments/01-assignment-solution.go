/*
rewrite to take advantange of go concurrency
also, print the prime numbers in the main() function
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := findPrimes(1000, 2000)
	for _, primeNo := range primes {
		fmt.Println("Prime :", primeNo)
	}
}

func findPrimes(start, end int) []int {
	var mutex sync.Mutex
	var primes []int
	wg := sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				mutex.Lock()
				{
					primes = append(primes, no)
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
