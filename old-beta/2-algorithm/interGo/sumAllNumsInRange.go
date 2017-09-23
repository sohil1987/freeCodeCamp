package main

import "fmt"

func sumAllNumsInRange() {
	var sol = make([]int, 0)
	for _, v := range test15 {
		//fmt.Println(v.data)
		sol = append(sol, getSumInRange(v.data))
	}
	for i, v := range sol {
		fmt.Println(i, v)
	}
}

func getSumInRange(nums [2]int) int {
	var sum = 0
	var min, max int
	if nums[0] < nums[1] {
		min = nums[0]
		max = nums[1]
	} else {
		min = nums[1]
		max = nums[0]
	}
	for i := min; i <= max; i++ {
		sum += i
	}
	return sum
}

var test15 = []struct {
	data [2]int
}{
	{[2]int{1, 4}},
	{[2]int{4, 1}},
	{[2]int{5, 10}},
	{[2]int{10, 5}},
}
