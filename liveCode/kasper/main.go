package main

import "fmt"

// func main() {
// 	ch1 := make(chan int, 1)
// 	ch1 <- 1
// 	for i := 0; i < 100; i++ {
// 		select {
// 		case <-ch1:
// 		case <-ch1:
// 			ch1 <- 1
// 		default:
// 			fmt.Println(i)
// 			return
// 		}

// 	}
// }

func main() {
	msgchan := make(chan string, 1)
	close(msgchan)
	stopchan := make(chan struct{})
	close(stopchan)
	select {
	case msgchan <- "msg":
		fmt.Println("msg sent")
	case <-stopchan:
		fmt.Println("stop signal received")
	}
}