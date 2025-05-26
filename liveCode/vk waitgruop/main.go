package main

import "sync"

type LimitedGroup struct {
    wg       sync.WaitGroup
    mx       sync.Mutex
    chanPool chan struct{}
    errJob   error
}

func NewLimitedGroup(size int) *LimitedGroup {
    chanPool := make(chan struct{}, size)
    
    return &LimitedGroup{
        wg:       sync.WaitGroup{},
        mx:       sync.Mutex{},
        chanPool: chanPool,
        errJob:   nil,
    }
}

// Do запускает задачу или блокируется, если уже слишком много запущенных задач.
func (lg *LimitedGroup) Do(job func() error) {
    lg.chanPool <- struct{}{}
    lg.wg.Add(1)
    go func() {
        defer lg.wg.Done()
        err := job()
        lg.mx.Lock()
        if err != nil {
            lg.errJob = err
        }
        lg.mx.Unlock()
        <-lg.chanPool
    }()
}

// TryDo пытается запустить задачу. Если превышен лимит количества горутин, то возвращаем false
func (lg *LimitedGroup) TryDo(job func() error) bool {
    select {
    case lg.chanPool <- struct{}{}:
        lg.wg.Add(1)
        go func() {
            defer lg.wg.Done()
            err := job()
            lg.mx.Lock()
            if err != nil {
                lg.errJob = err
            }
            lg.mx.Unlock()
            <-lg.chanPool
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
