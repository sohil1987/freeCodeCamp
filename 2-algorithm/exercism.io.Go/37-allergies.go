package main

import (
	"fmt"
	"math"
)

func allergies() {
	a := alergies{ //map[int]string{
		1:   "eggs",
		2:   "peanuts",
		4:   "shellfish",
		8:   "strawberries",
		16:  "tomatoes",
		32:  "chocolate",
		64:  "pollen",
		128: "cats",
	}
	for _, v := range test37 {
		fmt.Sprintln(v.input, a)
		list := getAllergies(int(v.input), a)
		fmt.Println(v.input, " --> ", list)
	}
}

func getAllergies(num int, a alergies) []string {
	res := make([]string, 0)
	arrayPos := getPosArr(num)
	fmt.Sprintln(arrayPos)
	return res
}

func getPosArr(num int) []int {
	arr := []int{}
	index := 0
	val := func() int { return int(math.Pow(2, float64(index))) }
	found := false
	for !found {
		//fmt.Println("arr", index, val(), num)
		if val() > num {
			found = true
		} else {
			arr = append(arr, val())
		}
		index++
	}
	return arr
}

func getPosMap(num int) int {
	i := 0
	for i < 32 {
		if math.Pow(2, float64(i)) >= float64(num) {
			return i
		}
		i++
	}
	return i
}

type alergies map[int]string

var test37 = []struct {
	expected []string
	input    uint
}{ /*
		{[]string{}, 0},
		{[]string{"eggs"}, 1},
		{[]string{"peanuts"}, 2},
		{[]string{"strawberries"}, 8},
		{[]string{"eggs", "peanuts"}, 3},
		{[]string{"eggs", "shellfish"}, 5},
		{[]string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 248},
		{[]string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 255},*/
	{[]string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 509},
}
