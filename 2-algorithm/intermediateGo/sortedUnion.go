package main

import (
	"fmt"
	s "strconv"
	"strings"
)

func uniteUnique(arr ...[]int) []string {
	var result string
	var aux []string
	//fmt.Println(arr)
	for _, v := range arr {
		//fmt.Println(v)
		for _, d := range v {
			if !strings.Contains(result, s.Itoa(d)) {
				aux = append(aux, s.Itoa(d))
				result = strings.Join(aux[:], "")
			}
		}
	}
	return strings.Split(result, "")
}

func sortedUnion() {
	fmt.Println(uniteUnique([]int{1, 3, 2}, []int{5, 2, 1, 4}, []int{2, 1}))
	fmt.Println(uniteUnique([]int{1, 3, 2}, []int{1, 5}, []int{2, 4}))
	fmt.Println(uniteUnique([]int{1, 2, 3}, []int{5, 2, 1}))
	fmt.Println(uniteUnique([]int{1, 2, 3}, []int{5, 2, 1, 4}, []int{2, 1}, []int{6, 7, 8}))
}
