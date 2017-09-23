package main

import "fmt"

func sortedUnion() {
	var sol = make([][]int, 0)
	for _, v := range test13 {
		var aux = make([]int, 0)
		//fmt.Println(i, v.data, len(v.data))
		for _, t := range v.data {
			//fmt.Println(t)
			for _, x := range t {
				//fmt.Println(x)
				if !sliceContainsInt(x, aux) {
					aux = append(aux, x)
				}
			}
		}
		sol = append(sol, aux)
	}
	for _, v := range sol {
		fmt.Println(v)
	}
}

func sliceContainsInt(num int, slice []int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

var test13 = []struct {
	data [][]int
}{
	{[][]int{
		[]int{1, 3, 2}, []int{5, 2, 1, 4}, []int{2, 1},
	}},
	{[][]int{
		[]int{1, 3, 2}, []int{1, 5}, []int{2, 4},
	}},
	{[][]int{
		[]int{1, 2, 3}, []int{5, 2, 1},
	}},
	{[][]int{
		[]int{1, 2, 3}, []int{5, 2, 1, 4}, []int{2, 1}, []int{6, 7, 8},
	}},
}
