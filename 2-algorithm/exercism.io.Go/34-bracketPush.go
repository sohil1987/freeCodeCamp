package main

import (
	"fmt"
)

func bracketPush() {
	for _, t := range test34 {
		isOk := iterateBrackets(t.input)
		fmt.Printf("%-12s is %-5t ,is OK? %t\n", t.input, isOk, isOk == t.expected)
	}
}

func iterateBrackets(str string) bool {
	status := map[int]int{123: 0, 125: 0, 91: 0, 93: 0, 40: 0, 41: 0}
	help := [][]int{{123, 125}, {91, 93}, {40, 41}}
	var duo int
	for _, v := range str {
		if v == 123 || v == 125 {
			duo = 0
		} else if v == 91 || v == 93 {
			duo = 1
		} else {
			duo = 2
		}
		pos := int(v)
		//fmt.Println(status)
		if pos == help[duo][0] {
			if status[help[duo][0]] >= status[help[duo][1]] {
				status[pos]++
			} else if status[help[duo][0]] < status[help[duo][1]] {
				status[pos]++
			}
		} else if pos == help[duo][1] {
			if status[help[duo][0]] > status[help[duo][1]] {
				status[pos]++
			} else if status[help[duo][0]] <= status[help[duo][1]] {
				return false
			}
		}
		//fmt.Println(status)
	}
	if status[123] != status[125] || status[91] != status[93] ||
		status[40] != status[41] {
		return false
	}
	return true
}

var test34 = []struct {
	input    string
	expected bool
}{
	{
		input:    "",
		expected: true,
	},
	{
		input:    "{}",
		expected: true,
	},
	{
		input:    "{{",
		expected: false,
	},
	{
		input:    "}{",
		expected: false,
	},
	{
		input:    "{}[]",
		expected: true,
	},
	{
		input:    "{[]}",
		expected: true,
	},
	{
		input:    "{[}]",
		expected: false,
	},
	{
		input:    "{[)][]}",
		expected: false,
	},
	{
		input:    "{[]([()])}",
		expected: true,
	},
}
