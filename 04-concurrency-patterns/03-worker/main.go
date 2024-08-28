package main

import (
	"fmt"
	"math/rand"
	"time"
	"worker-demo/worker"
)

type MyWork struct {
	Id int
}

/* worker.Work interface implementation */
func (myWork MyWork) Task() {
	fmt.Printf("task [%d] started...\n", myWork.Id)
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	fmt.Printf("task [%d] completed...\n", myWork.Id)
}

func main() {
	/* Approach:1 - sequential execution of the tasks (NOT performant) */
	/*
		for i := 1; i <= 20; i++ {
			myWork := MyWork{Id: i}
			myWork.Task()
		}
		fmt.Println("Done")
	*/

	// Approach:2 - 1 goroutine per task (NOT resource efficient)
	/*
		wg := sync.WaitGroup{}
		for i := 1; i <= 20; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				myWork := MyWork{Id: id}
				myWork.Task()
			}(i)
		}
		wg.Wait()
		fmt.Println("Done")
	*/

	/* Approach:3 */

	w := worker.New(5)
	for i := 1; i <= 20; i++ {
		w.Add(MyWork{Id: i})
	}
	fmt.Println("All tasks are assigned")
	w.Shutdown() // wait for all the assigned tasks to complete and DO NOT accept any more tasks
	fmt.Println("Done")
}
