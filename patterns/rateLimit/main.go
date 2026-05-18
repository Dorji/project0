package main

import (
	"fmt"
)

// Воркер пул через каналы
func main() {
	chanForWorker := make(chan string)
	for i := 0; i < 3; i++ {
		go worker(chanForWorker, i) // запускаем воркера
	}

	months := []string{
		"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	for _, month := range months {
		chanForWorker <- month //кидаем задачи
	}
	close(chanForWorker) // обязательно закрываем иначе не дождёмся выполнения, закрывается (close()) когда нет данных внутри канала
}

func worker(in <-chan string, th int) {
	for v := range in {
		fmt.Println("thread ", th, "get ", v)
	}
}
