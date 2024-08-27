/* channel behavior */
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100 // (2.NB)
	}()
	data := <-ch // (1.B) (3.UB)
	fmt.Println(data)
}
