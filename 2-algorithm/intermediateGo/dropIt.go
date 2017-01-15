package main

import (
	"fmt"
)

type function func(int) bool

func dropElements(arr []int, f function) []int {
	var res []int
	for _, v := range arr {
		if f(v) {
			fmt.Println("true", v)
			res = append(res, v)
		}
	}
	return res
}

func dropIt() {
	fmt.Println(dropElements([]int{1, 2, 3}, func(n int) bool { return n < 3 }))
	fmt.Println(dropElements([]int{1, 2, 3, 4}, func(n int) bool { return n >= 3 }))
}
