package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

// Задача1: сформировать массив из N разных случайных элементов с помощью math/rand
// Задача2: проверить строку на паллидром (одинаково читается с разных концов)
// Задача3: заменить заданный символ в слайсе на 3 других (заданных), inplace (предполагаем что места хватит)
// Задача4: дано слайс чисел, заменить каждое на произведение остальных
// ======= все решать оптимально
func main(){
	fmt.Println("Задача1")
	fmt.Println(createRandArray(5))

	fmt.Println("Задача2")
	fmt.Println(palyndrome("abcd","dcba"))
	fmt.Println(palyndrome("abcd","dcbqqa"))
	fmt.Println(palyndrome("2abcd","dcba1"))
	fmt.Println(palyndrome("1abcd","dcba1"))

	fmt.Println("Задача3")
	arrInput:=strings.Split("123456abcd12345","")
	arrInput=append(arrInput,"","","" )
	fmt.Println(inputInSlice(arrInput,6,[3]string{"2","2","2"}))

	fmt.Println("Задача4")
	fmt.Println(multiplicationOfOthers([]int{1,2,3,4,5}))
}
// Задача1: сформировать массив из N разных случайных элементов с помощью math/rand
func createRandArray(n int) []interface{}{
	
	res:=make([]interface{}, n)

	for i:=0;i<n;i++{	typeOfValue:=rand.Intn(4)
		switch	typeOfValue{
			case 1:
				res[i]=rand.Int()
			case 2:  
				length:=rand.Intn(10)
				b := make([]rune, length)
				for i := range b {
					b[i] = rune(rand.Intn(utf8.MaxRune))
				}
				res[i]=string(b)  
			case 3:
				res[i]=rand.Float64()
			case 4:
				res[i]=bool((rand.Int()%2)>0)
		}
	}
	return res
}
// Задача2: проверить строку на паллидром (одинаково читается с разных концов)
func palyndrome(str1 string, str2 string) bool{
		arr1:=strings.Split(str1,"")
		arr2:=strings.Split(str2,"")
		if len(arr2)!=len(arr1) || (len(arr1)==0||len(arr2)==0){
			return false
		}
		for k,_ := range arr1{
			if arr1[k]!=arr2[len(arr2)-1-k]{
				return false
			}
		}
		return true
}

// Задача3: заменить заданный символ в слайсе на 3 других (заданных), inplace (предполагаем что места хватит)
func inputInSlice(slc []string,n int, val [3]string) []string {
	slc3:=make([]string,0)
	slc3 = append(slc3,slc[n+1:]...)
	fmt.Println(slc3)
	slc = append(slc[:n],val[0],val[1],val[2])
	slc =  append(slc[:n+3],slc3[:len(slc3)-3]...)
	return slc
}

// Задача4: дано слайс чисел, заменить каждое на произведение остальных
func multiplicationOfOthers([]int)[]int{
	res:=[]int{}
	for k,v :=range 

	return res
}