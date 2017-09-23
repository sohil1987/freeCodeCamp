package main

import (
	"fmt"
)

func nthPrime() {
	for _, t := range test51 {
		res := getPrimePos(t.n)
		fmt.Printf("%10d is in %-6d position, is OK? %5t\n", t.n, res, t.order == res)
	}
}

func getPrimePos(n int) int {
	pos := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			pos++
		}
	}
	return pos
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var test51 = []struct {
	order int
	n     int
	ok    bool
}{
	{1, 2, true},
	{2, 3, true},
	{3, 5, true},
	{4, 7, true},
	{5, 11, true},
	{6, 13, true},
	{10001, 104743, true},
	{0, 0, false},
}
