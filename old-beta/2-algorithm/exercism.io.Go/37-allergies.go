package main

import (
	"fmt"
	"math"
	"reflect"
)

func allergies() {
	for _, t := range test37 {
		res := make([]string, 0)
		res = getAllergies(int(t.input))
		fmt.Printf("is OK? %5t  %v\n", reflect.DeepEqual(t.expected, res), t.expected)
	}
}

func getAllergies(n int) []string {
	var res = make([]string, 0)
	exp := len(a) - 1
	for exp >= 0 {
		val := int(math.Pow(2, float64(exp)))
		if n >= val {
			res = append(res, a[val])
			n -= val
		}
		exp--
	}
	return res
}

type alergies map[int]string

var a = alergies{ //map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

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
