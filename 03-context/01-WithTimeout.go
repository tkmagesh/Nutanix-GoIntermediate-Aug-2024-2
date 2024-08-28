package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// send the cancellation signal
		cancel()
	}()
	go genNos(cancelCtx, wg)
	wg.Wait()
	fmt.Println("Done")
}

func genNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		}

	}
}
