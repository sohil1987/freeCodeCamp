package main

import (
	"fmt"
	"strconv"
)

func raindrops() {
	for _, v := range test06 {
		var res = ""
		if v.input%3 == 0 {
			res += "Pling"
		}
		if v.input%5 == 0 {
			res += "Plang"
		}
		if v.input%7 == 0 {
			res += "Plong"
		}
		if res == "" {
			res = strconv.Itoa(v.input)
		}
		fmt.Printf("Solution = %-20s , is OK ? %t\n", res, v.expected == res)
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
	{8, "8"},
	{9, "Pling"},
	{10, "Plang"},
	{14, "Plong"},
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{27, "Pling"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
	{3125, "Plang"},
}
