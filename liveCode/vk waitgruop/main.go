package main


type LimitedGroup struct {
    wg sync.WaitGroup
    mx sync.Mutex
    chanPool chan struct{}
    errJob error
    
}


func NewLimitedGroup(size int) *LimitedGroup {
    errSlc := make([]error,0)
    chanPool := make(chan struct{},size)
    
    
    
    return &LimitedGroup{
        wg: sync.WaitGroup{},
        mx: sync.Mutex{},
        chanPool: chanPool,
        errJob: nil,
    }
}


// Do запускает задачу или блокируется, если уже слишком много запущенных задач.
func (lg *LimitedGroup) Do(job func() error {}) {
    chanPool<-struct{}
    lg.wg.Add(1)
    go func(){
        defer lg.wg.Done()
            errJob:=job()
            lg.mx.Lock()
            if lg.errJob!=nil{
                lg.errJob=err
            }
            lg.mx.Unlock()
        <-chanPool
    }()
}

// TryDo пытается запустить задачу. Если превышен лимит количества горутин, то возвращаем false
func (lg *LimitedGroup) TryDo(job func() error {}) bool {
        select{
            case chanPool<-struct{}:
                 lg.wg.Add(1)
                     go func(){
                    defer lg.wg.Done()
                     err:=job()
                     lg.mx.Lock()
                     if lg.errJob!=nil{
                         lg.errJob=err
                     }
                     lg.mx.Unlock()
                    <-chanPool
                  }()
                  return true
            default:
                return false
        }
    
}

// Wait блокируется пока все горутины не завершили выполнение
func (lg *LimitedGroup) Wait() error {
    defer close(lg.chanPool)
         lg.wg.Wait()
         return lg.errJob
}




package main

import "log"
import "runtime"

var a string
var done bool

func setup() {
    done = true
	a = "hello, world"
	if done {
		log.Println(len(a))
	}
}

func main() {
	go setup()

	for !done {
		runtime.Gosched()
	}
	log.Println(a)
}

// 11
// "hello, world"

package main

type (
    Person interface {
        Action() error
    }

    Worker struct {
        ...
    }
)

func Foo(person Person) error {
    return person.Action()
}

func Bar(persons []Person) error {
    ...
}

func main() {
    worker := &Worker{}
    _ := Foo(worker)

    workers := []Worker{&Worker{}, &Worker{}, ..., &Worker{}}
    _ := Bar(workers)
}













