package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func diffArray(s1, s2 interface{}) []string {
	var aux1 []int
	var res, a1, a2 []string
	//fmt.Println(reflect.TypeOf(s1))
	if reflect.TypeOf(s1) == reflect.TypeOf(aux1) {
		//fmt.Println(s1, s2, "son []int")
		a1, a2 = convertIntToString(s1.([]int), s2.([]int))
	} else {
		a1, a2 = s1.([]string), s2.([]string)
	}
	// aqui ya s1,s2 son []string
	for _, v1 := range a1 {
		var found = false
		for _, v2 := range a2 {
			//fmt.Println(v1, v2)
			if v1 == v2 {
				//fmt.Println("FOUND ", v1)
				found = true
			}
		}
		if !found {
			//fmt.Println(`NOT FOUND`, v1)
			res = append(res, v1)
		}
	}
	for _, v2 := range a2 {
		var found = false
		for _, v1 := range a1 {
			//fmt.Println(v1, v2)
			if v1 == v2 {
				//fmt.Println("FOUND ", v2)
				found = true
			}
		}
		if !found {
			//fmt.Println(`NOT FOUND`, v2)
			res = append(res, v2)
		}
	}
	//fmt.Println(res)
	return res
}

func convertIntToString(s1, s2 []int) ([]string, []string) {
	//fmt.Println(s1, s2)
	var res1, res2 []string
	for _, v := range s1 {
		res1 = append(res1, strconv.Itoa(v))
	}
	for _, v := range s2 {
		res2 = append(res2, strconv.Itoa(v))
	}
	//fmt.Println(res1, res2)
	return res1, res2
}

func diffTwoArrays() {
	fmt.Println(diffArray([]int{1, 2, 3, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(diffArray([]int{1, 2, 3, 5, 6, 7}, []int{1, 2, 3, 4, 5, 8}))
	fmt.Println(diffArray([]string{"diorite", "andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(diffArray([]string{"andesite", "grass", "dirt", "pink wool", "dead shrub"}, []string{"diorite", "andesite", "grass", "dirt", "dead shrub"}))
	fmt.Println(diffArray([]string{"andesite", "grass", "dirt", "dead shrub"}, []string{"andesite", "grass", "dirt", "dead shrub"}))
}
