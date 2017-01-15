package main

import "fmt"

type function2 func(int) bool

func findElement(arr []int, f function2) int {
	var res int
	found := false
	i := 1
	for !found && i < len(arr) {
		//fmt.Println(i, arr[i])
		if f(arr[i]) {
			found = true
			res = arr[i]
		}
		i++
	}
	return res
}

func findersKeepers() {
	fmt.Println(findElement([]int{1, 3, 5, 8, 9, 10}, func(n int) bool { return n%2 == 0 }))
	fmt.Println(findElement([]int{1, 3, 5, 9}, func(n int) bool { return n%2 == 0 }))
}
