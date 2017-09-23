package main

import (
	"fmt"
	"math"
	"strconv"
)

func palindrome() {
	for _, test := range test31 {
		res := make([]int, 0)
		for i := test.min; i <= test.max; i++ {
			for j := i; j <= test.max; j++ {
				if isPalindromeProduct(i * j) {
					res = append(res, i*j)
				}
			}
		}
		fmt.Println(test.min, test.max, res)

	}
}

func isPalindromeProduct(num int) bool {
	s := strconv.Itoa(num)
	half := int(math.Floor(float64(len(s) / 2)))
	for i := 0; i <= half; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

var test31 = []struct {
	min, max int
}{
	{1, 9},
	{10, 99},
	{100, 999},
	{4, 10},
	{10, 4},
}
