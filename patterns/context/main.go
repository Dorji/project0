package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cancelWithTimeout()
	//cancelOnFirst()
}

func cancelOnFirst() {

	ctx, finish := context.WithCancel(context.Background()) // получаем функцию финиш
	result := make(chan int, 1)

	for i := 1; i < 10; i++ {
		go worker(ctx, i, result) //она передаётся в контекст
	}

	foundBy := <-result // Ждём результат
	fmt.Println("result found by", foundBy)
	finish() // и отключаем
	time.Sleep(time.Second)
}
func cancelWithTimeout() {
	workTime := 5000 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	result := make(chan int, 1)

	for i := 1; i < 10; i++ {
		go worker(ctx, i, result)
	}

	totalFound := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case foundBy := <-result:
			totalFound++
			fmt.Println("result found by", foundBy)
		}

	}
}

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}
