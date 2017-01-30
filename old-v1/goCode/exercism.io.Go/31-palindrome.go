package main

import "fmt"
import "strconv"
import "math"

func palindrome() {
	min := 100
	max := 999
	res := iterateRange(min, max)
	fmt.Println(min, max, res)
}

func isPalindromeProduct(num int) bool {
	s := strconv.Itoa(num)
	//fmt.Println("MIRANDO ", s)
	half := int(math.Floor(float64(len(s) / 2)))
	for i := 0; i <= half; i++ {
		//fmt.Println("kk", i, half, len(s))
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func iterateRange(min, max int) []int {
	var res []int
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			//fmt.Println(i, j)
			if isPalindromeProduct(i * j) {
				fmt.Println(`Bingo`, i, j)
				res = append(res, i*j)
			}
		}
	}
	return res
}
