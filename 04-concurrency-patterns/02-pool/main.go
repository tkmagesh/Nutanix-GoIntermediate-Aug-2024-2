package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pool-demo/pool"
	"sync"
	"time"
)

/*
	Resource = any object implements io.Closer interface
	Factory = function that creates an instance of the resource
*/

// Resource
type DBConnection struct {
	Id int
}

// io.Closer interface implementation
func (dbConnection DBConnection) Close() error {
	fmt.Printf("Resource [id=%d] is being discarded\n", dbConnection.Id)
	return nil
}

// resource factory
var idCounter int

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	dbConnection := DBConnection{Id: idCounter}
	return dbConnection, nil
}

func main() {
	/*
		create a Pool instance (with the pool size, factory function)

		Resource => any object that implements "close()" method (io.Closer interface)

		When a resource is Acquire()d
			the pool will check if it has any resources
			if yes, return the resource from the pool
			else create a new resource (using the factory) and return the resource
			Make sure that when a resource is acquired, it not acquired again until it is released

		When a resource is Release()d to the pool
			the pool will check if it is full
			if yes, then discard the resource (after 'closing' the resource)
			else, keep the resource to serve future requests (maintain the resource in the pool)

		When Close()ing the pool
			prevent anymore aqcuisition of the resources
			make sure all the resources are closed and discarded

		Notes:
			The number of requests for the resources CAN exceed the pool size
			When a resource is acquired by the client, the same resource SHOULD NOT be given to another client until the resource is released back to pool


		Functions:
			New(poolSize, factory) => pool instance

		Methods (pool)
			Acquire() => resource
			Release(resource)
			Close()

	*/

	p, err := pool.New(5 /* poolSize */, DBConnectionFactory)
	if err != nil {
		log.Fatalln(err)
	}
	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for c := 1; c <= clientCount; c++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(c)
	}
	wg.Wait()
	fmt.Println("Batch - 1 completed.  Press ENTER to continue...")

	fmt.Scanln()
	fmt.Println("Batch - 2 Starts (To verify if the resources are returned from the pool).")
	wg.Add(10)
	for c := 21; c <= 30; c++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(c)
	}
	wg.Wait()
	p.Close()
}

func doWork(client int, p *pool.Pool) {
	connResource, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("worker [id=%d]: Acquired resource [id=%d]\n", client, connResource.(DBConnection).Id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) //simulation of using the resource
	fmt.Printf("worker [id=%d]: Releasing resource [id=%d]\n", client, connResource.(DBConnection).Id)
	p.Release(connResource)
}
