//// Находим максимальное четное число

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//func main() {
//	var max int
//	wg:= &sync.WaitGroup{}
//	mx:=&sync.Mutex{}
//
//	for i := 1000; i > 0; i-- {
//		wg.Add(1)
//		go func(mx *sync.Mutex ,wg *sync.WaitGroup,i int,max *int) {
//			defer wg.Done()
//			mx.Lock()
//			if i%2 == 0 && i > *max {
//				*max = i
//			}
//			mx.Unlock()
//		}(mx,wg,i,&max)
//	}
//	wg.Wait()
//	fmt.Printf("Maximum is %d", max)
//}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(50)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc() int64 {
	ctx := context.Background()
	ctx, cncl := context.WithTimeout(ctx, 1*time.Second)
	ch := make(chan int64, 5)

	go func(ctx context.Context, ch chan int64) {
		val := unpredictableFunc()
		select {
		case <-ctx.Done():
			ch <- -1
		case ch <- val:
		}
	}(ctx, ch)
	var res int64
	select {
	case res = <-ch:
		cncl()
		return res
	case <-ctx.Done():
		cncl()
		return res
	}

}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())

}
