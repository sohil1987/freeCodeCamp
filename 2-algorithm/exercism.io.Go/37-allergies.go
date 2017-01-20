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
		list := getAllergies(int(v.input), a)
		fmt.Println(v.input, " --> ", list)
	}
}

func getAllergies(num int, a alergies) []string {
	res := make([]string, 0)
	arrayPos := getPosArr(num)
	arrayPos = reverseSliceInt(arrayPos)
	for _, v := range arrayPos {
		if num >= v {
			if _, ok := a[v]; ok {
				//fmt.Println(`HOLA `, v, a[v])
				res = append(res, a[v])
			}
			num -= v
		}
	}
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

func reverseSliceInt(reverse []int) []int {
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return reverse
}

type alergies map[int]string

var test37 = []struct {
	expected []string
	input    uint
}{
	{[]string{}, 0},
	{[]string{"eggs"}, 1},
	{[]string{"peanuts"}, 2},
	{[]string{"strawberries"}, 8},
	{[]string{"eggs", "peanuts"}, 3},
	{[]string{"eggs", "shellfish"}, 5},
	{[]string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 248},
	{[]string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 255},
	{[]string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 509},
}
