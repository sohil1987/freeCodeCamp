package main

import (
	"fmt"
	"sort"
)

func smallestCommons(arr []int) int {
	sort.Ints(arr)
	min := arr[0]
	max := arr[1]
	var result = max
	var test []int
	for i := min; i <= max; i++ {
		test = append(test, i)
	}
	var found = false
	for !found && result < 2000000000 {
		//fmt.Println(result)
		if isOk(result, test) {
			return result
		}
		result++
	}
	return 0
}

func isOk(n int, test []int) bool {
	for _, v := range test {
		if n%v != 0 {
			return false
		}
	}
	return true
}

func smallCommonMultiple() {
	fmt.Println(smallestCommons([]int{1, 5}))
	fmt.Println(smallestCommons([]int{5, 1}))
	fmt.Println(smallestCommons([]int{1, 13}))
	fmt.Println(smallestCommons([]int{23, 18}))
}
