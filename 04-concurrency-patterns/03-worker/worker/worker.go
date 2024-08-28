package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	workQueue chan Work
	wg        sync.WaitGroup
}

func (w *Worker) Add(wk Work) {
	w.workQueue <- wk
}

func (w *Worker) Shutdown() {
	close(w.workQueue)
	w.wg.Wait()
}

func New(workerCount int) *Worker {
	worker := &Worker{
		workQueue: make(chan Work),
		wg:        sync.WaitGroup{},
	}
	for i := 1; i <= workerCount; i++ {
		worker.wg.Add(1)
		go func(id int) {
			fmt.Printf("Worker %d started\n", id)
			defer worker.wg.Done()
			for work := range worker.workQueue {
				work.Task()
			}
		}(i)
	}
	return worker
}
