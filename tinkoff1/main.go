package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	mu := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(group *sync.WaitGroup, i int) {
			defer group.Done()

			mu.Lock()
			ch <- fmt.Sprintf("Goroutine %d", i)
			mu.Unlock()
		}(&wg, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

LOOP:
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				break LOOP
			}
			fmt.Println(s)
		}
	}

}
