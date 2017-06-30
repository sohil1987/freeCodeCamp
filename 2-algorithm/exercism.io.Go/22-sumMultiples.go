package main

import (
	"fmt"
)

func sumMultiples() {
	for i, test := range test22 {
		sum := getSumMultiples(test.divisors, test.limit)
		fmt.Printf("%-2d up to %-6d == %-10d,is OK? %t\n", i, test.limit, sum, sum == test.sum)
	}
}

func getSumMultiples(divs []int, limit int) int {
	if len(divs) == 0 || limit < 1 {
		return 0
	}
	list := make([]int, 0)
	sum := 0
	for i := 1; i < limit; i++ {
		for _, m := range divs {
			if i%m == 0 && !sliceContainsInt(i, list) {
				sum += i
				list = append(list, i)
			}
		}
	}
	return sum
}

func sliceContainsInt(num int, slice []int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

var test22 = []struct {
	divisors []int
	limit    int
	sum      int
}{
	{[]int{3, 5}, 1, 0},             // multiples of 3 or 5 up to 1
	{[]int{3, 5}, 4, 3},             // multiples of 3 or 5 up to 4
	{[]int{3, 5}, 10, 23},           // multiples of 3 or 5 up to 10
	{[]int{3, 5}, 100, 2318},        // multiples of 3 or 5 up to 100
	{[]int{3, 5}, 1000, 233168},     // multiples of 3 or 5 up to 1000
	{[]int{7, 13, 17}, 20, 51},      // multiples of 7, 13 or 17 up to 20
	{[]int{4, 6}, 15, 30},           // multiples of 4 or 6 up to 15
	{[]int{5, 6, 8}, 150, 4419},     // multiples of 5, 6 or 8 up to 150
	{[]int{5, 25}, 51, 275},         // multiples of 5 or 25 up to 51
	{[]int{43, 47}, 10000, 2203160}, // multiples of 43 or 47 up to 10000
	{[]int{1}, 100, 4950},           // multiples of 1 up to 100
	{[]int{}, 10000, 0},             // multiples of an empty list up to 10000
}
