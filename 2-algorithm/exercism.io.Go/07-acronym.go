package main

import (
	"fmt"
	"strings"
)

func acronym() {
	fmt.Println("acronym")
	for i, v := range test07 {
		var result string
		v.input = strings.Replace(v.input, "-", " ", -1)
		words := strings.Split(v.input, " ")
		for _, w := range words {
			result += strings.ToUpper(string(w[0]))
		}
		fmt.Println(i, v.input, result)
	}
}

var test07 = []struct {
	input    string
	expected string
}{
	{"Portable Network Graphics", "PNG"},
	{"HyperText Markup Language", "HTML"},
	{"Ruby on Rails", "ROR"},
	{"PHP: Hypertext Preprocessor", "PHP"},
	{"First In, First Out", "FIFO"},
	{"Complementary metal-oxide semiconductor", "CMOS"},
}
