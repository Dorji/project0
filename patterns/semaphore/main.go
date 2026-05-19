package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Semaphore struct {
	mx         *sync.Mutex
	wg         *sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	maxWorkers chan struct{}
	closed     bool
	errors     []error
}

func NewSemaphore(maxWorkerCount int) *Semaphore {
	ctx, cancel := context.WithCancel(context.Background())
	maxWork := make(chan struct{}, maxWorkerCount)
	return &Semaphore{
		mx:         &sync.Mutex{},
		wg:         &sync.WaitGroup{},
		ctx:        ctx,
		cancel:     cancel,
		maxWorkers: maxWork,
	}
}

func (s *Semaphore) Do(f func() error) {

	s.wg.Go(func() {
		s.maxWorkers <- struct{}{}
		defer func() { <-s.maxWorkers }()
		err := f()
		if err != nil {
			s.cancel()
			s.errors = append(s.errors, err)
		}
	})
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
	close(s.maxWorkers)
}

func main() {
	sem := NewSemaphore(2)

	fnc := func(val int) error {
		time.Sleep((time.Duration(rand.Intn(3))) * time.Second)
		fmt.Println(val)
		return nil
	}

	for i := range 15 {
		sem.Do(func() error {
			return fnc(i)
		})
	}

	sem.Wait()
	fmt.Println("work ended!")
}
