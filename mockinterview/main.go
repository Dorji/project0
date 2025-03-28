package main

import (
	"fmt"
	"strings"
)

// Задача 1: сформировать массив из N разных случайных элементов с помощью math/rand
// Задача 2: проверить строку на паллидром (одинаково читается с разных концов)
// Задача 3: заменить заданный символ в слайсе на 3 других (заданных), inplace (предполагаем что места хватит)
// Задача 4: дано слайс чисел, заменить каждое на произведение остальных
// ======= все решать оптимально
// 1) тут ты что-то напутал, у тебя есть одна функция ...rand()  -она тебе генерит одно число (можно взять библиотечную), нужно выполнить ВСЕ условия задачи
// 2) на входе один массив его и проверяют на перевертываемость
// не все кейсы рассмотрены
// 3 - я идею вообще не понял
// кстати есть testify - для тестов
// 4) делать надо in place
// без дополнительных выделений памяти, то есть входной массив будет выходным
func main() {
	fmt.Println("Задача 2")
	fmt.Println(palyndrome("abcd", "dcba"))
	fmt.Println(palyndrome("abcd", "dcbqqa"))
	fmt.Println(palyndrome("2abcd", "dcba1"))
	fmt.Println(palyndrome("1abcd", "dcba1"))

	fmt.Println("Задача 3")
	arrInput := strings.Split("123456abcd12345", "")
	arrInput = append(arrInput, "", "", "")
	fmt.Println(inputInSlice(arrInput, 6, [3]string{"2", "2", "2"}))

	fmt.Println("Задача 4")
	fmt.Println(multiplicationOfOthers([]int{1, 2, 3, 4, 5}))
	fmt.Println(multiplicationOfOthers([]int{1, 2, 3, 4, 5, 0}))
	fmt.Println(multiplicationOfOthers([]int{1, 0, 3, 4, 5, 0}))
	fmt.Println(multiplicationOfOthers([]int{1, 0, 3, 4, 5, 6}))
	fmt.Println(multiplicationOfOthers([]int{1, 2, 3, 4, 5, 6}))

}

// Задача 2: проверить строку на паллидром (одинаково читается с разных концов)
func palyndrome(str1 string, str2 string) bool {
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	if len(arr2) != len(arr1) || (len(arr1) == 0 || len(arr2) == 0) {
		return false
	}
	for k, _ := range arr1 {
		if arr1[k] != arr2[len(arr2)-1-k] {
			return false
		}
	}
	return true
}

// Задача 3: заменить заданный символ в слайсе на 3 других (заданных), inplace (предполагаем что места хватит)
func inputInSlice(slc []string, n int, val [3]string) []string {
	slc3 := make([]string, 0)
	slc3 = append(slc3, slc[n+1:]...)
	fmt.Println(slc3)
	slc = append(slc[:n], val[0], val[1], val[2])
	slc = append(slc[:n+3], slc3[:len(slc3)-3]...)
	return slc
}

// Задача 4: дано слайс чисел, заменить каждое на произведение остальных
func multiplicationOfOthers(slc []int) []int {
	res := make([]int, len(slc))
	val := 1
	zeroElem := -1
	for k, _ := range slc {
		if slc[k] == 0 {
			if zeroElem == -1 {
				zeroElem = k
			} else {
				return make([]int, len(slc), len(slc))
			}
		} else {
			val = slc[k] * val
		}
	}
	for k, _ := range slc {
		if k != zeroElem {
			res[k] = val / slc[k]
		} else {
			res[k] = 0
		}
	}
	return res
}
