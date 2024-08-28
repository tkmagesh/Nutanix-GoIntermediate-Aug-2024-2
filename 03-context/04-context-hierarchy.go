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
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

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
	fmt.Println("[genNos] root-key :", ctx.Value("root-key"))

	dupRootKeyCtx := context.WithValue(ctx, "root-key", "new-root-key")
	genNosCtx := context.WithValue(dupRootKeyCtx, "nos-key", "nos-value")

	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	evenCtx, cancel := context.WithTimeout(genNosCtx, 5*time.Second)
	defer cancel()
	go genEvenNos(evenCtx, wg1)

	wg1.Add(1)
	oddCtx, cancel := context.WithTimeout(genNosCtx, 7*time.Second)
	defer cancel()
	go genOddNos(oddCtx, wg1)

LOOP:
	for i := 0; ; i += 1 {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("[genNos] programmatic cancellation signal received")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[genNos] cancellation signal received by timeout")
			}
			break LOOP
		default:
			fmt.Println("No:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg1.Wait()

}

func genEvenNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[genEvenNos] root-key :", ctx.Value("root-key"))
	fmt.Println("[genEvenNos] nos-key :", ctx.Value("nos-key"))
LOOP:
	for i := 0; ; i += 2 {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("[genEvenNos] programmatic cancellation signal received")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[genEvenNos] cancellation signal received by timeout")
			}
			break LOOP
		default:
			fmt.Println("Even No:", i)
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func genOddNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[genOddNos] root-key :", ctx.Value("root-key"))
	fmt.Println("[genOddNos] nos-key :", ctx.Value("nos-key"))
LOOP:
	for i := 1; ; i += 2 {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("[genOddNos] programmatic cancellation signal received")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[genOddNos] cancellation signal received by timeout")
			}
			break LOOP
		default:
			fmt.Println("Odd No:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
