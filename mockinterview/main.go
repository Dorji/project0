package main 

import(
	"fmt"
	"strings"
)
// Задача1: сформировать массив из N разных случайных элементов с помощью math/rand
// Задача2: проверить строку на паллидром (одинаково читается с разных концов)
// Задача3: заменить заданный символ в слайсе на 3 других (заданных), inplace (предполагаем что места хватит)
// Задача4: дано слайс чисел, заменить каждое на произведение остальных
// ======= все решать оптимально
func main(){
	fmt.Println(palyndrome("abcd","dcba"))
	fmt.Println(palyndrome("abcd","dcbqqa"))
	fmt.Println(palyndrome("2abcd","dcba1"))
	fmt.Println(palyndrome("1abcd","dcba1"))
}

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