package main

import (
	"fmt"
)

func perfectsNums() {
	for _, v := range test33 {
		//fmt.Println(v.w)
		numType := getTypeNumber(v.w)
		switch numType {
		case "perfect":
			fmt.Println(v.w, " --> PERFECT")
		case "abundant":
			fmt.Println(v.w, " --> ABUNDANT")
		case "deficient":
			fmt.Println(v.w, " --> DEFICIENT")
		}
	}
}

func getTypeNumber(num int) string {
	var sum int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	if sum > num {
		return "abundant"
	} else if sum < num {
		return "deficient"
	}
	return "perfect"
}

var test33 = []struct {
	w int
}{
	{1},
	{13},
	{12},
	{6},
	{28},
	{496},
	{8128},
}
