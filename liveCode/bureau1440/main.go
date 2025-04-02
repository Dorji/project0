package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) {
	start := time.Now()
	ch := make(chan int64, 1)

	go func() {
		ch <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		duration := time.Since(start)
		fmt.Printf("Function execution timed out after %v\n", duration)
		return 0, ctx.Err()
	case res := <-ch:
		duration := time.Since(start)
		fmt.Printf("Function completed successfully in %v\n", duration)
		return res, nil
	}
}

func main() {
	fmt.Println("started")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := predictableFunc(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}