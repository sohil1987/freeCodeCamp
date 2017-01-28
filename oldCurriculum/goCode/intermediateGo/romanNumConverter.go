package main

import (
	"fmt"
	"math"
)

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

func romanNumConverter() {
	fmt.Println(convertToRoman(36))
	fmt.Println(convertToRoman(68))
	fmt.Println(convertToRoman(97))
	fmt.Println(convertToRoman(1000))
	fmt.Println(convertToRoman(3999))
}
