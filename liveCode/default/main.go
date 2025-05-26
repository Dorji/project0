

package main

import "log"
import "runtime"

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
	worker := &Worker{}
    _ := Foo(worker)

    workers := []Worker{&Worker{}, &Worker{}, ..., &Worker{}}
    _ := Bar(workers)

	go setup()

	for !done {
		runtime.Gosched()
	}
	log.Println(a)


}














