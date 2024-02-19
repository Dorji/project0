package main

import (
	"fmt"
	"time"
)

func main() {
}
func ticker() {

	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)
	return
}
func timers() {
	// при 1 выполнится таймаут, при 3 - выполнится операция
	timer := time.NewTimer(3 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("simple timer")
	case <-time.After(time.Minute):
		//Не убирается пока не выполнится
		fmt.Println("time.After happened")
	case result := <-longSQLQuery():
		// освобождет ресурс
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result:", result)
	}

}
func longSQLQuery() chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()
	return ch
}
