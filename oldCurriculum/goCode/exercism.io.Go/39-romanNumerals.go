package main

import (
	"fmt"
	"math"
)

func romanNumerals() {
	for _, v := range test39 {
		res := convertToRoman(v.arabic)
		fmt.Println(res)
	}
}

func convertToRoman(n int) string {
	var res string
	hundreds := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	units := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	t := int(math.Floor(float64(n / 1000)))
	h := int(math.Floor(float64(n % 1000 / 100)))
	d := int(math.Floor(float64(n % 100 / 10)))
	u := int(math.Floor(float64(n % 10)))
	for i := 0; i < t; i++ {
		res += "M"
	}
	if h > 0 {
		res += hundreds[h-1]
	}
	if d > 0 {
		res += tens[d-1]
	}
	if u > 0 {
		res += units[u-1]
	}
	return res
}

var test39 = []struct {
	arabic   int
	roman    string
	hasError bool
}{
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
