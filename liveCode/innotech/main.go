/*

Слить N каналов в один
Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.
Мы хотим, чтобы результат работы функции выглядел примерно так:
for num := range joinChannels(a, b, c) {
     fmt.Println(num)
}

*/
package main

import (
    // "context"
    "fmt"
    "sync"
)

func main() {

    a := make(chan int)
    b := make(chan int)
    c := make(chan int)

    go func() {
        defer close(a)
        for i := 0; i < 5; i++ {
            a <- i
        }
    }()

    go func() {
        defer close(b)
        for i := 5; i < 10; i++ {
            b <- i
        }
    }()

    go func() {
        defer close(c)
        for i := 10; i < 15; i++ {
            c <- i
        }
    }()
    
    out := joinChannels(a, b, c)
    for num := range out {
        fmt.Println(num)
    }
}

func joinChannels(chs ...<- chan int) chan int{
     res:=make(chan int) 
     
    wg:=&sync.WaitGroup{} 
    
        for _,v := range chs{
            wg.Add(1)
            go func(){
               defer wg.Done()
               for{
                    if val,ok := <-v; ok{
                         res<-val
                    }else{
                         return
                    }
               }
            }()
        } 
    
    go func(){
        wg.Wait()
      close(res)
    }()
    
    return res
}





















