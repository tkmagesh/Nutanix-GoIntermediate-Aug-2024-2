package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("Pool closed")

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	sync.Mutex
}

func (p *Pool) Acquire() (io.Closer, error) {
	p.Lock()
	defer p.Unlock()
	select {
	case resource, isOpen := <-p.resources:
		if !isOpen {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return resource, nil
	default:
		fmt.Println("Acquiring from the factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.Lock()
	defer p.Unlock()
	select {
	case p.resources <- resource:
		fmt.Println("Releasing the resource to the pool")
		return nil
	default:
		fmt.Println("pool full. discarding the resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.Lock()
	defer p.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}

func New(poolSize int, factory func() (io.Closer, error)) (*Pool, error) {
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		closed:    false,
	}, nil
}
