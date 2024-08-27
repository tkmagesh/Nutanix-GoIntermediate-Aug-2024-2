/*
rewrite to take advantange of go concurrency
also, print the prime numbers in the main() function
*/
package main

import "fmt"

func main() {
	findPrimes(1000, 2000)
	fmt.Println("Done!")
}

func findPrimes(start, end int) {
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Println("Prime :", no)
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
