package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	months := []string{
		"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	for _, month := range months {
		wg.Add(1) // перед запуском горутины а не в ней потому-что
		// горутина может запуститься позже чем цикл, и в wg.wait будет 0 и всё остановится
		// т.е. мы потеряем данные
		go startWorker(month, wg) //кидаем дерьмо лопатами параллельно
	}
	wg.Wait()

}

func startWorker(cnt string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(cnt, " worked")
}
