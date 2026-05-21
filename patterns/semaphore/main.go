package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Semaphore struct {
	mu         sync.Mutex
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	maxWorkers chan struct{}
	closed     bool
	errors     []error
}

func NewSemaphore(maxWorkerCount int) *Semaphore {
	ctx, cancel := context.WithCancel(context.Background())
	return &Semaphore{
		maxWorkers: make(chan struct{}, maxWorkerCount),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (s *Semaphore) Do(f func() error) error {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return fmt.Errorf("семафор остановлен")
	}
	s.wg.Add(1) // Добавляем ДО запуска горутины
	s.mu.Unlock()

	go func() {
		defer s.wg.Done()

		// Ждем слот или остановку семафора
		select {
		case s.maxWorkers <- struct{}{}:
			defer func() { <-s.maxWorkers }()
		case <-s.ctx.Done():
			return
		}

		// Выполняем задачу, если контекст ещё активен
		if s.ctx.Err() != nil {
			return
		}

		if err := f(); err != nil {
			s.mu.Lock()
			s.errors = append(s.errors, err)
			s.mu.Unlock()
		}
	}()

	return nil
}

func (s *Semaphore) Shutdown() []error {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return s.errors
	}
	s.closed = true
	s.mu.Unlock()

	s.wg.Wait()
	close(s.maxWorkers)
	s.cancel()

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.errors
}

func main() {
	sem := NewSemaphore(2)

	fnc := func(val int) error {
		sleepTime := time.Duration(rand.Intn(3)) * time.Second
		fmt.Printf("Задача %d началась (спим %v)\n", val, sleepTime)
		time.Sleep(sleepTime)
		fmt.Printf("Задача %d завершена\n", val)
		return nil
	}

	for i := range 15 {
		id := i // Захватываем переменную
		err := sem.Do(func() error {
			return fnc(id)
		})
		if err != nil {
			fmt.Printf("Ошибка запуска: %v\n", err)
		}
	}

	errors := sem.Shutdown()
	
	if len(errors) > 0 {
		fmt.Println("\nОшибки:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	}
	
	fmt.Println("Все задачи завершены!")
}