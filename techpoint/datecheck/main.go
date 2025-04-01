package main

import (
    "fmt"
    "math/rand"
    "time"
    "context"
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
func predictableFunc(ctx context.Context) (int64,error) {
    ch:= make(chan int64,1)
    defer close(ch)
    
    
    go func(){
         ch<-unpredictableFunc()
    }()

    select {
        case <-ctx.Done():
            return 0,fmt.Errorf("MyError")
        case res:=<-ch:
            return res,nil
    }
    
     
     return 0,fmt.Errorf("MyError")
}

func main() {
    fmt.Println("started")
    
    ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

    fmt.Println(predictableFunc(ctx))
}