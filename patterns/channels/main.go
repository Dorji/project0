package main

import "fmt"

func main() {
	channels()
	channels1()

}
func channels() {
	ch1 := make(chan int)    // небуферизированный канал
	ch2 := make(chan int, 1) // буферизированный канал

	go func(in chan int) {
		val := <-in
		fmt.Println("go get from chan", val)
		fmt.Println("go after read from chan")
	}(ch1)
	go func(in chan int) {
		val := <-in
		fmt.Println("go get from chan", val)
		fmt.Println("go after read from chan")
	}(ch2)

	ch1 <- 42
	//ch1<-100500 // fatal error deadlock! limitless blocking because nobody listening
	ch2 <- 100500
	ch2 <- 500100 // no deadlock, can be not complete

	fmt.Println("Main:after put in chan")
}

func channels1() {
	in := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {

			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out) // LOOP will be stopped here
		fmt.Println("generator finish")
	}(in)
	//LOOP:
	for i := range in {
		fmt.Println("\tget", i)
	}

}
