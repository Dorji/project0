package main

import "fmt"

func main() {
	select2()
}

func select2() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(cancelCh <-chan struct{}, dataCh chan<- int) {
		val := 0
		for {
			//В селекте если есть несколько совпадений выбирается случайное
			select {
			case <-cancelCh: // срабатывает закрытие
				return
			case dataCh <- val: //бесконечно пишем в канал
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{} // Посылаем в канал закрытия сигнал о закрытии
			break
		}
	}

}

func select1() {

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3

LOOP:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("ch1 val", v1)
		case v2 := <-ch2:
			fmt.Println("ch2 val", v2)
		default:
			break LOOP

		}
	}
}
