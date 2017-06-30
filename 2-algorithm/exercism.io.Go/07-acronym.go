package main

import (
	"fmt"
	"strings"
)

func acronym() {
	for _, v := range test07 {
		aux := strings.Replace(v.input, "-", " ", -1)
		words := strings.Split(aux, " ")
		var res = ""
		for _, t := range words {
			res += strings.ToUpper(string(t[0]))
		}
		fmt.Printf("Solution = %-20s , is OK ? %t\n", res, v.expected == res)
	}

}

type t07 struct {
	input    string
	expected string
}

var test07 = []t07{
	{input: "Portable Network Graphics", expected: "PNG"},
	{input: "Ruby on Rails", expected: "ROR"},
	{input: "First In, First Out", expected: "FIFO"},
	{input: "PHP: Hypertext Preprocessor", expected: "PHP"},
	{input: "GNU Image Manipulation Program", expected: "GIMP"},
	{input: "Complementary metal-oxide semiconductor", expected: "CMOS"},
}

var test07b = []struct {
	input    string
	expected string
}{
	{"Portable Network Graphics", "PNG"},
	{input: "Ruby on Rails", expected: "ROR"},
	{input: "First In, First Out", expected: "FIFO"},
	{input: "PHP: Hypertext Preprocessor", expected: "PHP"},
	{input: "GNU Image Manipulation Program", expected: "GIMP"},
	{input: "Complementary metal-oxide semiconductor", expected: "CMOS"},
}
