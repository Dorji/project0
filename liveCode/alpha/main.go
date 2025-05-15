package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	// fmt.Println(randUniqSlice(5))
	WorkerPool(5)
}
//1 Написать код для выдачи массива уникальных случайных чисел
func randUniqSlice(n int)[]int{
	tmpMap:=make(map[int]int,n)
	res:=make([]int,0,0)
	LOOP:
	for {
		val:=rand.Int()
		tmpMap[val]=val
		if len(tmpMap)==n{
			break LOOP
		}
	}
	for k:= range tmpMap{
		res=append(res,tmpMap[k])
	}

	return res
}

//2 Нужно написать HTTP-сервер, который обрабатывает GET-запросы и возвращает список всех файлов в указанной директории


//3 Реализовать worker pool. Есть 10 задач (функций), каждая засыпает на 1 сек и выводит номер воркера, который эту задачу исполнил. Количество воркеров задается при запуске.

func WorkerPool(num int){
	chanWork:=make(chan *Work ,0)
	
	for i:=0;i<num;i++{
		go worker( num, chanWork )
	}

	
	for i:=0;i<10;i++{
		chanWork<-&Work{ number:i}
	}
	
	close(chanWork)

}

type Work struct{
	number int
}

func (w *Work)StartWork(){
	time.Sleep(time.Second)
	fmt.Println(w.number)
}
func worker(  num int, chanWork chan *Work){
	for val:=range chanWork{
		val.StartWork()
	}
		
}