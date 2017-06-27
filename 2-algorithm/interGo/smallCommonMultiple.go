package main

import "fmt"

func smallCommonMultiple() {
	var sol = make([]int, 0)
	for _, v := range test12 {
		var min, max int
		if v.data[0] <= v.data[1] {
			min = v.data[0]
			max = v.data[1]
		} else {
			min = v.data[1]
			max = v.data[0]
		}
		var worstSol = getWorstSol(min, max)
		//fmt.Printf("Min %-5d Max %-5d WorstSol %-5d\n", min, max, worstSol)
		index := max
		var found = false
		for index <= worstSol && found == false {
			if iFoundIt(min, max, index) {
				found = true
				sol = append(sol, index)
			}
			index++
		}
	}
	for _, v := range sol {
		fmt.Println(v)
	}
}

func iFoundIt(min, max, index int) bool {
	for i := min; i <= max; i++ {
		if index%i != 0 {
			return false
		}
	}
	return true
}

func getWorstSol(min, max int) int {
	var worst = 1
	for i := min; i <= max; i++ {
		worst = worst * i
	}
	return worst
}

var test12 = []struct {
	data [2]int
}{
	{[2]int{1, 5}},
	{[2]int{5, 1}},
	{[2]int{2, 10}},
	{[2]int{1, 13}},
	{[2]int{23, 18}},
}
