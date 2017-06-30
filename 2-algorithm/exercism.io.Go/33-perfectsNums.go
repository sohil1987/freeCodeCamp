package main

import (
	"fmt"
)

func perfectsNums() {
	for _, test := range test33 {
		typeNumber := getTypeNumber(int(test.input))
		fmt.Printf("%-5d is %-12s ,is OK? %t\n", test.input, typeNumber, typeNumber == test.expected)
	}
}

func getTypeNumber(num int) string {
	types := map[int]string{0: "Deficient", 1: "Abundant", 2: "Perfect"}
	option := 0
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	if sum == num {
		option = 2
	} else if sum > num {
		option = 1
	} else {
		option = 0
	}
	return types[option]
}

var test33 = []struct {
	input    uint64
	expected string
}{
	{1, "Deficient"},
	{13, "Deficient"},
	{12, "Abundant"},
	{6, "Perfect"},
	{28, "Perfect"},
	{496, "Perfect"},
	{8128, "Perfect"},
}
