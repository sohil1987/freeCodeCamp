package main

import (
	"fmt"
	"strconv"
)

func raindrops() {
	for _, v := range test06 {
		var result string
		//fmt.Println(i, v.input)
		if v.input%3 == 0 {
			result += "Pling"
		}
		if v.input%5 == 0 {
			result += "Plang"
		}
		if v.input%7 == 0 {
			result += "Plong"
		}
		if result != "" {
			fmt.Println(v.input, " --> ", result)
		} else {
			fmt.Println(v.input, " --> ", strconv.Itoa(v.input))
		}
	}
}

var test06 = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
	{6, "Pling"},
	{9, "Pling"},
	{10, "Plang"},
	{14, "Plong"},
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
}
