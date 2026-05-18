package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	WorkerPool(5)
}
func WorkerPool(numWorkers int) {
	// Вариант 1: Буферизированный канал
	chanWork := make(chan *Work, 10) // буфер на все задачи
	
	// Вариант 2: Отправка в отдельной горутине
	// chanWork := make(chan *Work)
	
	wg := &sync.WaitGroup{}
	
	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, wg, chanWork) // i+1 чтобы нумерация была с 1
	}
	
	// Отправляем задачи
	for i := 0; i < 10; i++ {
		chanWork <- &Work{number: i + 1}
	}
	
	close(chanWork) // Сигнал воркерам что задач больше не будет
	wg.Wait()       // Ждем завершения всех воркеров
}

type Work struct {
	number int
}

func (w *Work) StartWork(workerID int) {
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d выполнил задачу %d\n", workerID, w.number)
}

func worker(id int, wg *sync.WaitGroup, chanWork <-chan *Work) {
	defer wg.Done()
	for work := range chanWork {
		work.StartWork(id)
	}
}