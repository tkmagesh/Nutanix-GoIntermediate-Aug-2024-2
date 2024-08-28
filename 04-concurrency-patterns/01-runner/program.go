package main

import (
	"fmt"
	"os"
	"runner-demo/runner"
	"time"
)

func main() {
	/*
		initialize the runner with a timeout
		add multiple tasks to the runner
		Start the runner
		if all the tasks are completed within the given time, report "Success"
		if the tasks are not completed within the given time, report "timeout"
		exit if the execution is interrupted by an OS interrupt
	*/
	fmt.Printf("Process %d started....\n", os.Getpid())
	timeout := 15 * time.Second
	r := runner.New(timeout)

	r.Add(createTask(5))
	r.Add(createTask(7))
	r.Add(createTask(2))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("task execution timed out")
		case runner.ErrInterrupt:
			fmt.Println("OS interrupt received. Shutting Down....")
		default:
			fmt.Println("unknown error :", err)
		}
	} else {
		fmt.Println("All tasks are executed within the given time")
	}
}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Printf("Start Processing task # %d\n", id)
		time.Sleep(time.Duration(t) * time.Second)
		fmt.Printf("Completed Processing task # %d\n", id)
	}
}
