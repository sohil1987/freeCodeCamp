package main

import (
	"fmt"
	"math"
)

func romanNumConverter() {
	var sol = make([]string, 0)
	for _, v := range test10 {
		sol = append(sol, getRomanNum(v.data))
	}
	for i, v := range sol {
		fmt.Println(i, v)
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

var test10 = []struct {
	data int
}{
	{36},
	{68},
	{697},
	{1000},
	{3999},
}
