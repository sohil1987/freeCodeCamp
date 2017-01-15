package main

import (
	"fmt"
)

func sumPrimes(n int) int {
	var result int
	for i := 2; i <= n; i++ {
		//fmt.Println(i)
		if isPrime(i) {
			result += i
		}
	}
	return result
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	//fmt.Printf("%d is Prime", n)
	return true
}

func sumAllPrimes() {
	fmt.Println(sumPrimes(10))
	fmt.Println(sumPrimes(977))
}
