package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := program()
	if err != nil {
		return
	}
	for k := range res {
		fmt.Println(res[k])
	}

}
func program() ([]string, error) {
	res := make([]string, 0)
	var total int
	_, err := fmt.Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("no total")
	}
	slc := make([][]string, total, total)
	for i := 0; i < total; i++ {
		value := make([]string, 10, 10)
		slc[i] = value
		for iteration := 0; iteration < 10; iteration++ {
			scannedValue := ""
			_, err := fmt.Scan(&scannedValue)
			slc[i][iteration] = scannedValue
			if err != nil {
				return nil, fmt.Errorf("error in line")
			}
		}
	}

	//etalonMap :=make(map[int]int)
	referenceMap := map[int]int{
		4: 1,
		3: 2,
		2: 3,
		1: 4,
	}

	iterationMap := map[int]int{
		4: referenceMap[4],
		3: referenceMap[3],
		2: referenceMap[2],
		1: referenceMap[1],
	}
	for k := range slc {
		for key := range slc[k] {
			keyNum, err := strconv.Atoi(slc[k][key])
			if err != nil {
				return nil, fmt.Errorf("cannot parsing")
			}
			if _, ok := iterationMap[keyNum]; ok {
				iterationMap[keyNum] = iterationMap[keyNum] - 1

			}
		}

		cnt := 0

	LOOP:
		for j := range iterationMap {
			if iterationMap[j] < 0 {
				res = append(res, "NO")
				cnt = 1
				break LOOP
			}
		}
		if cnt != 1 {
			res = append(res, "YES")
		}
		cnt = 0
		iterationMap = map[int]int{
			4: referenceMap[4],
			3: referenceMap[3],
			2: referenceMap[2],
			1: referenceMap[1],
		}
	}
	return res, nil
}
