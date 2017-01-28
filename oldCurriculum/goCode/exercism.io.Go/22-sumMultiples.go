package main

import "fmt"

func sumMultiples() {
	for _, v := range test22 {
		multiples := []int{}
		//fmt.Println(v.limit, v.divisors)
		limit := v.limit
		for _, multiple := range v.divisors {
			multiples = calcMultiples(multiple, limit, multiples)
		}
		sum := sumM(multiples)
		fmt.Println("TOTAL --> ", sum)
	}
}

func sumM(m []int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func calcMultiples(num, lim int, m []int) []int {
	//fmt.Println(m)
	for i := 2; i < lim; i++ {
		if !sliceContainsInt(m, i) {
			if i%num == 0 {
				m = append(m, i)
			}
		}
	}
	//fmt.Println(m)
	return m
}

func sliceContainsInt(m []int, n int) bool {
	for _, v := range m {
		if v == n {
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
	{[]int{3, 5}, 1, 0},
	{[]int{3, 5}, 4, 3},
	{[]int{3, 5}, 10, 23},
	{[]int{3, 5}, 100, 2318},
	{[]int{3, 5}, 1000, 233168},
	{[]int{7, 13, 17}, 20, 51},
	{[]int{43, 47}, 10000, 2203160},
	{[]int{5, 10, 12}, 10000, 13331672},
	{[]int{1, 1}, 10000, 49995000},
	{[]int{}, 10000, 0},
}
