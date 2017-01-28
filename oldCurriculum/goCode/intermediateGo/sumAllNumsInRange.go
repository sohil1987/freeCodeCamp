package main

import (
	"fmt"
)

func sumAll(arr [2]int) int {
	var sum int
	min := arr[0]
	max := arr[1]
	if min > max {
		aux := min
		min = max
		max = aux
	}
	for i := min; i <= max; i++ {
		sum += i
	}
	return sum
}

func sumAllNumsIntRange() {
	fmt.Println(sumAll([2]int{1, 4}))
	fmt.Println(sumAll([2]int{4, 1}))
	fmt.Println(sumAll([2]int{5, 10}))
	fmt.Println(sumAll([2]int{10, 5}))
}
