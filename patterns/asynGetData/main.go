package main

import (
	"fmt"
	"time"
)

// Асинхронное получение данных для отложенного получения данных и парралельных работ
func main() {
	for i := 0; i < 3; i++ {
		getPage()
	}

}
func getPage() {
	resultCh := getComments()
	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")
	//log.Fatal("sdsds") в этом случае здесь будет утечка горутин
	//
	commentsData := <-resultCh
	//commentsData=<-resultCh
	fmt.Println("main goroutine:", commentsData)
}

func getComments() chan string {
	// надо использовать буферизированный канал
	result := make(chan string, 2)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second) //какая-то работа
		fmt.Println("async operation ready, return comments")
		out <- "загрузка комментариев"
		time.Sleep(1 * time.Second)
		out <- "32 комментария"
	}(result)

	return result
}
