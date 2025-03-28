package main

import (
	"fmt"
	"testing"
		"math/rand"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	fmt.Println("Задача 1")
	assert.Equal(t, 5, len(createRandArray(5)))
	fmt.Println(createRandArray(5))

	// fmt.Println("Задача2")
	// assert.Equal(t, palyndrome("abcd", "dcba"), true)
	// assert.Equal(t, palyndrome("abcd", "dcba1"), false)

}

// Задача1: сформировать массив из N разных случайных элементов с помощью math/rand
func createRandArray(n int) []int {

	res := make([]int, 0)

	tmp:=make(map[int]int,0)
	prevLen:=0

	for i := 0; i < n; {
		val:=rand.Intn(10)
		fmt.Println(val)
		tmp[val]=val
		if len(tmp)>prevLen{
			i++
		}
	}
	for k,_:=range tmp{
		res=append(res,k)
	}

	return res
}
