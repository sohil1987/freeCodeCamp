package main

import (
	"fmt"
	"math"
	"reflect"
)

func romanNumerals() {
	for _, t := range test39 {
		res := getRomanNum(t.arabic)
		fmt.Printf("is OK? %5t  %v\n", reflect.DeepEqual(t.roman, res), t.roman)
	}
}

func getRomanNum(num int) (sol string) {
	hundreds := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	units := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	t := int(math.Floor(float64(num / 1000)))
	h := int(math.Floor(float64(num % 1000 / 100)))
	d := int(math.Floor(float64(num % 100 / 10)))
	u := int(math.Floor(float64(num % 10)))
	for i := 0; i < t; i++ {
		sol += "M"
	}
	if h > 0 {
		sol += hundreds[h-1]
	}
	if d > 0 {
		sol += tens[d-1]
	}
	if u > 0 {
		sol += units[u-1]
	}
	return sol
}

type t39 struct {
	arabic   int
	roman    string
	hasError bool
}

var test39 = []t39{
	{1, "I", false},
	{2, "II", false},
	{3, "III", false},
	{4, "IV", false},
	{5, "V", false},
	{6, "VI", false},
	{9, "IX", false},
	{27, "XXVII", false},
	{48, "XLVIII", false},
	{59, "LIX", false},
	{93, "XCIII", false},
	{141, "CXLI", false},
	{163, "CLXIII", false},
	{402, "CDII", false},
	{575, "DLXXV", false},
	{911, "CMXI", false},
	{1024, "MXXIV", false},
	{3000, "MMM", false},
}
