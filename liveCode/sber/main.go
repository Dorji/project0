// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"sync"
)

type AsyncPool struct {
	maxChan chan struct{}
	wg      sync.WaitGroup
	mx      sync.Mutex
}

func NewAsyncPool(size int) *AsyncPool {

	return &AsyncPool{
		maxChan: make(chan struct{}),
		wg:      sync.WaitGroup{},
		mx:      sync.Mutex{},
	}
}

func (a *AsyncPool) TryDo(work func()) error {
	a.maxChan<- struct{}{}
	a.wg.Add(1)
	go func ()  {
		defer a.wg.Done()
		err:= work()
		a.mx.Lock()
		<-a.maxChan
	}()
}

func (a *AsyncPool) Wait() {
	a.wg.Wait()
}
