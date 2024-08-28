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

	// to share data thorugh context
	valCtx := context.WithValue(rootCtx, "start", 100)

	// time base context
	timeoutCtx, cancel := context.WithTimeout(valCtx, 10*time.Second)

	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// send the cancellation signal
		cancel()
	}()
	go genNos(timeoutCtx, wg)
	wg.Wait()
	fmt.Println("Done")
}

func genNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	start := ctx.Value("start").(int)
LOOP:
	for i := start; ; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("programmatic cancellation signal received")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("cancellation signal received by timeout")
			}
			break LOOP
		default:
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		}

	}
}
