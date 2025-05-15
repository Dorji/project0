package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	WorkerPool(5)
} 
//Реализовать worker pool. Есть 10 задач (функций), каждая засыпает на 1 сек и выводит номер воркера, который эту задачу исполнил. Количество воркеров задается при запуске.

func WorkerPool(num int){
	chanWork:=make(chan *Work ,0)
	// ctx,cancel:=context.WithCancel(context.Background())
	wg:= &sync.WaitGroup{}
	for i:=0;i<num;i++{
		wg.Add(1)
		go worker(wg, num, chanWork )
	}

	
	for i:=0;i<10;i++{
		chanWork<-&Work{ number:i}
	}
	
	close(chanWork)
	wg.Wait()
}

type Work struct{
	number int
}

func (w *Work)StartWork(){
	time.Sleep(time.Second)
	fmt.Println(w.number)
}
func worker( wg *sync.WaitGroup, num int, chanWork chan *Work){
	defer wg.Done()
	for val:=range chanWork{
		val.StartWork()
	}
		
}