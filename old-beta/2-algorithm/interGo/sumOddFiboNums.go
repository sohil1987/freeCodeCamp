package main

import "fmt"

func sumOddFiboNums() {
	var sol = make([]int, 0)
	for _, v := range test17 {
		//fmt.Println(v.data)
		sol = append(sol, getSumOddFiboNums(v.data))
	}
	for i, v := range sol {
		fmt.Printf("%-2d == %d \n", i, v)
	}
}

func getSumOddFiboNums(n int) int {
	if n == 1 {
		return 1
	}
	var sum = 2
	var fib = []int{1, 1}
	for i := 2; fib[i-1] < n; i++ {
		new := fib[i-2] + fib[i-1]
		fib = append(fib, new)
		if fib[i] <= n && fib[i]%2 != 0 {
			sum += fib[i]
		}
	}
	//fmt.Println(len(fib), "-->", fib)
	return sum
}

var test17 = []struct {
	data int
}{
	{1},
	{1000},
	{4000000},
	{4},
	{75024},
	{75025},
}
