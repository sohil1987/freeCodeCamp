package main

import (
	"fmt"
)

func sumFibs(limit int) int {
	var sum = 2
	var fib = []int{1, 1}
	for i := 2; fib[i-1] < limit; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
		if fib[i] > limit { // end, limit exceeded
			return sum
		} else if fib[i]%2 != 0 {
			sum += fib[i]
		}
	}
	//fmt.Println(fib)
	return sum
}

func sumOddFiboNums() {
	fmt.Println(sumFibs(1))
	fmt.Println(sumFibs(1000))
	fmt.Println(sumFibs(4000000))
	fmt.Println(sumFibs(4))
	fmt.Println(sumFibs(75024))
	fmt.Println(sumFibs(75025))

}
