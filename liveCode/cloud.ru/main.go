package main

import (
	"fmt"
	"sync"
	"time"
)

type Work struct{}

func (w *Work) startWork() {
	time.Sleep(5 * time.Second)
}

type WorkHandler struct {
	chWork chan *Work       // канал для передачи задач воркерам
	chPool chan struct{}    // семафор для ограничения параллельных задач
	queue  []*Work          // очередь задач
	maxSize int             // максимальный размер очереди
	mtx     sync.Mutex      // мьютекс для безопасного доступа к очереди
}

func NewHandler(maxSize int, n int) (*WorkHandler, error) {
	if n == 0 {
		return nil, fmt.Errorf("no pool size")
	}

	if maxSize == 0 {
		return nil, fmt.Errorf("no maxSize")
	}

	return &WorkHandler{
		chWork:  make(chan *Work),
		chPool:  make(chan struct{}, n), // буферизированный канал как семафор
		queue:   make([]*Work, 0, maxSize),
		maxSize: maxSize,
	}, nil
}

func (wh *WorkHandler) Start() {
	// Запускаем горутину для обработки очереди
	go wh.processQueue()
	
	// Инициализируем пул воркеров
	for i := 0; i < cap(wh.chPool); i++ {
		wh.chPool <- struct{}{} // заполняем пул доступными слотами
		go wh.worker()
	}
}

func (wh *WorkHandler) processQueue() {
	for {
		wh.mtx.Lock()
		if len(wh.queue) > 0 {
			work := wh.queue[0]
			wh.queue = wh.queue[1:]
			wh.mtx.Unlock()
			
			wh.chWork <- work // отправляем задачу воркеру
		} else {
			wh.mtx.Unlock()
			time.Sleep(100 * time.Millisecond) // небольшая пауза при пустой очереди
		}
	}
}

func (wh *WorkHandler) Add(w *Work) error {
	wh.mtx.Lock()
	defer wh.mtx.Unlock()
	
	if len(wh.queue) >= wh.maxSize {
		return fmt.Errorf("queue is full")
	}
	
	wh.queue = append(wh.queue, w)
	return nil
}

func (wh *WorkHandler) worker() {
	for {
		// Ждем свободного слота в пуле
		<-wh.chPool
		
		// Получаем задачу
		work := <-wh.chWork
		
		// Выполняем работу
		work.startWork()
		
		// Возвращаем слот в пул
		wh.chPool <- struct{}{}
	}
}