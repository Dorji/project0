package main

import (
	//rftn "coursera/reflection"
	//evnlb "evenLab"
	"fmt"
	"math"
)

func main() {
	fmt.Println(m_round(4.72))
	fmt.Println(m_round(2.35))
	fmt.Println(m_round(2.5))
	fmt.Println(m_round(6))
	fmt.Println(m_round(3.13))
}

func m_round(v float64) float64 {
	floor := math.Floor(v)
	dr := v - floor
	var res float64 = 0
	if dr > 0.5 && dr < 0.75 {
		res = 0.5
	} else if dr > 0.5 && dr > 0.75 {

		res = 1
	} else if dr < 0.5 && dr > 0.25 {

		res = 0.5
	} else if dr < 0.5 && dr < 0.25 {

		res = 0
	} else if dr == 0.5 {
		res = 0.5
	}
	return res + floor
}
