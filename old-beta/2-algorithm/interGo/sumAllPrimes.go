package main

import "fmt"

func sumAllPrimes() {
	var sol = make([]int, 0)
	for _, v := range test16 {
		//fmt.Println(v.data)
		var sum = 0
		for index := 2; index <= v.data; index++ {
			if isPrime(index) {
				sum += index
			}
		}
		sol = append(sol, sum)
	}
	for i, v := range sol {
		fmt.Printf("%-5d == %d \n", i, v)
	}
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

var test16 = []struct {
	data int
}{
	{10},
	{977},
}
