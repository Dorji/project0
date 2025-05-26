package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Work struct{}

func (w *Work) startWork() {
	time.Sleep(5 * time.Second)
}

type WorkHandler struct {
	chWork   chan *Work       // канал для задач
	chPool   chan struct{}    // семафор пула
	queue    []*Work          // очередь задач
	maxSize  int              // макс. размер очереди
	mtx      sync.Mutex       // для очереди
	wg       sync.WaitGroup   // для ожидания завершения
	stopChan chan struct{}    // для остановки
	cond     *sync.Cond       // для уведомлений о новых задачах
}

func NewHandler(maxSize, poolSize int) (*WorkHandler, error) {
	if poolSize == 0 {
		return nil, errors.New("pool size must be > 0")
	}
	if maxSize == 0 {
		return nil, errors.New("max queue size must be > 0")
	}

	wh := &WorkHandler{
		chWork:   make(chan *Work),
		chPool:   make(chan struct{}, poolSize),
		queue:    make([]*Work, 0, maxSize),
		maxSize:  maxSize,
		stopChan: make(chan struct{}),
	}
	wh.cond = sync.NewCond(&wh.mtx)

	// Заполняем пул
	for i := 0; i < poolSize; i++ {
		wh.chPool <- struct{}{}
	}

	return wh, nil
}

func (wh *WorkHandler) Start() {
	wh.wg.Add(1)
	go wh.processQueue()

	// Запускаем воркеры
	for i := 0; i < cap(wh.chPool); i++ {
		go wh.worker()
	}
}

func (wh *WorkHandler) Stop() {
	close(wh.stopChan)
	wh.wg.Wait()
	close(wh.chWork)
}

func (wh *WorkHandler) processQueue() {
	defer wh.wg.Done()

	for {
		select {
		case <-wh.stopChan:
			return
		default:
			wh.mtx.Lock()
			for len(wh.queue) == 0 {
				wh.cond.Wait()
				select {
				case <-wh.stopChan:
					wh.mtx.Unlock()
					return
				default:
				}
			}

			work := wh.queue[0]
			wh.queue = wh.queue[1:]
			wh.mtx.Unlock()

			select {
			case wh.chWork <- work:
			case <-wh.stopChan:
				return
			}
		}
	}
}

func (wh *WorkHandler) Add(w *Work) error {
	wh.mtx.Lock()
	defer wh.mtx.Unlock()

	if len(wh.queue) >= wh.maxSize {
		return errors.New("queue is full")
	}

	wh.queue = append(wh.queue, w)
	wh.cond.Signal() // Уведомляем о новой задаче
	return nil
}

func (wh *WorkHandler) worker() {
	for work := range wh.chWork {
		func() {
			// Берем слот только когда есть работа
			<-wh.chPool
			defer func() {
				if r := recover(); r != nil {
					log.Printf("worker panic: %v", r)
				}
				wh.chPool <- struct{}{} // Всегда возвращаем слот
			}()

			work.startWork()
		}()
	}
}

// Пример использования
func main() {
	handler, err := NewHandler(100, 5)
	if err != nil {
		log.Fatal(err)
	}

	handler.Start()

	// Добавляем задачи
	for i := 0; i < 10; i++ {
		if err := handler.Add(&Work{}); err != nil {
			log.Printf("failed to add work: %v", err)
		}
	}

	// Даем время на обработку
	time.Sleep(3 * time.Second)

	// Graceful shutdown
	handler.Stop()
}