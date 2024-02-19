package main

import "fmt"

type MyErr struct{}

func (m MyErr) Error() string {
	return "error"
}

func main() {
	fmt.Println(returnError() == nil)      // f t
	fmt.Println(returnErrorPtr() == nil)   // f t
	fmt.Println(returnMyError() == nil)    // f<nil>
	fmt.Println(returnMyErrorPtr() == nil) // f
	fmt.Println(returnMyErrorNil() == nil) // f(nil)
}

func returnError() error {
	var err error
	return err
}

func returnErrorPtr() *error {
	var err *error
	return err
}

func returnMyError() error {
	var err MyErr
	return err
}

func returnMyErrorPtr() error {
	var err *MyErr
	return err
}

func returnMyErrorNil() *MyErr {
	return nil
}
