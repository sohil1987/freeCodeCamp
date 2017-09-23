package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func diffTwoArrays() {
	var sol = make([][]string, 0)
	for _, v := range test05 {
		var aux []int
		var d1, d2 []string
		if reflect.TypeOf(v.data1) == reflect.TypeOf(aux) {
			d1 = convertIntToString(v.data1.([]int))
			d2 = convertIntToString(v.data2.([]int))
		} else {
			d1, d2 = v.data1.([]string), v.data2.([]string)
		}
		//fmt.Println(d1, d2)
		sol = append(sol, getDiffTwoArrays(d1, d2))
	}
	for _, v := range sol {
		fmt.Println(v)
	}
}

func getDiffTwoArrays(d1, d2 []string) (sol []string) {
	var base = make([]string, len(d1))
	copy(base, d1)
	base = append(base, d2...)
	for _, v := range base {
		if sliceContainsString(v, d1) && sliceContainsString(v, d2) {
		} else {
			sol = append(sol, v)
		}
	}
	return sol
}

func convertIntToString(origin []int) (done []string) {
	for _, v := range origin {
		done = append(done, strconv.Itoa(v))
	}
	return done
}

func sliceContainsString(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

var test05 = []struct {
	data1 interface{}
	data2 interface{}
}{
	{
		[]int{1, 2, 3, 5},
		[]int{1, 2, 3, 4, 5},
	}, {
		[]int{1, 2, 3, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 8},
	}, {
		[]string{"diorite", "andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"},
	}, {
		[]string{"andesite", "grass", "dirt", "pink wool", "dead shrub"},
		[]string{"diorite", "andesite", "grass", "dirt", "dead shrub"},
	}, {
		[]string{"andesite", "grass", "dirt", "dead shrub"},
		[]string{"andesite", "grass", "dirt", "dead shrub"},
	},
}
