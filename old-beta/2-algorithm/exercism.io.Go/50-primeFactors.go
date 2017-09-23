package main

import (
	"fmt"
	"reflect"
)

func primeFactors() {
	for _, t := range test50 {
		res := getPrimeFactors(t.input)
		fmt.Printf("%12d gets %v is OK? %5t\n", t.input, res, reflect.DeepEqual(t.expected, res))
	}
}

func getPrimeFactors(n int64) []int64 {
	var res []int64
	if n == 1 {
		return res
	}
	trys := 0
	for n > 1 && trys < 5 {
		next := getNextDivisor(n)
		res = append(res, next)
		n /= next
		trys++
	}
	return res
}

func getNextDivisor(n int64) int64 {
	var next int64
	for next = 2; next <= n; next++ {
		if n%next == 0 {
			return next
		}
	}
	return next
}

var test50 = []struct {
	input    int64
	expected []int64
}{
	{1, []int64{}},
	{2, []int64{2}},
	{3, []int64{3}},
	{4, []int64{2, 2}},
	{6, []int64{2, 3}},
	{8, []int64{2, 2, 2}},
	{9, []int64{3, 3}},
	{27, []int64{3, 3, 3}},
	{625, []int64{5, 5, 5, 5}},
	{901255, []int64{5, 17, 23, 461}},
	{93819012551, []int64{11, 9539, 894119}},
}
