package main

import "fmt"

func main() {
	//non-buffered channel
	ch := make(chan int)

	//buffered channel (size 1)
	// ch := make(chan int, 1)

	//buffered channel (size 2)
	// ch := make(chan int, 2)

	//buffered channel (size 3)
	//ch := make(chan int, 3)

	go writeData(ch)

	fmt.Println("[@main] attempting to receive 10")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 10, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 20")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 20, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 30")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 30, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 40")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 40, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 50")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 50, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 60")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 60, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 70")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 70, len(ch) :", len(ch))

	fmt.Println("[@main] attempting to receive 80")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed receiving 80, len(ch) :", len(ch))

}

func writeData(ch chan int) {
	fmt.Println("	[@writeData] attempting to send 10")
	ch <- 10
	fmt.Println("	[@writeData] completed sending 10, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 20")
	ch <- 20
	fmt.Println("	[@writeData] completed sending 20, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 30")
	ch <- 30
	fmt.Println("	[@writeData] completed sending 30, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 40")
	ch <- 40
	fmt.Println("	[@writeData] completed sending 40, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 50")
	ch <- 50
	fmt.Println("	[@writeData] completed sending 50, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 60")
	ch <- 60
	fmt.Println("	[@writeData] completed sending 60, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 70")
	ch <- 70
	fmt.Println("	[@writeData] completed sending 70, len(ch) :", len(ch))

	fmt.Println("	[@writeData] attempting to send 80")
	ch <- 80
	fmt.Println("	[@writeData] completed sending 80, len(ch) :", len(ch))
}
