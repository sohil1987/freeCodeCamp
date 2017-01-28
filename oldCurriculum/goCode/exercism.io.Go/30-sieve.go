package main

import "fmt"

func sieve() {
	num := 1000
	arr := createArray(num)
	arr = deleteMultiples(arr)
	arr = cleanSolution(arr)
	fmt.Println(arr)
}

func cleanSolution(arr []int) []int {
	sol := make([]int, 0)
	for _, v := range arr {
		if v != 0 {
			sol = append(sol, v)
		}
	}
	fmt.Println(len(sol))
	return sol
}

func deleteMultiples(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	for i, v := range arr {
		if i > 1 {
			//fmt.Println(i, v)
			for j, val := range arr {
				//fmt.Println(j, val)
				if val > v && val%v == 0 {
					res[j] = 0
				}
			}
		}
	}
	return res
}

func createArray(num int) []int {
	arr := make([]int, 0)
	for i := 0; i <= num; i++ {
		if i == 1 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, i)
		}
	}
	return arr
}
