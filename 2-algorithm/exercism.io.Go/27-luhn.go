package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func luhn() {
	for _, test := range test27 {
		isValid := isValidLuhn(test.input)
		fmt.Printf("%-20s gets %t ,is OK? %t\n", test.input, isValid, isValid == test.ok)
	}
}

func isValidLuhn(s string) bool {
	clean := ""
	if string(s[0]) == " " {
		return false
	}
	s = strings.Replace(s, " ", "", -1)
	if string(s[0]) == "0" && len(s) == 1 {
		return false
	}
	for _, value := range s {
		if !unicode.IsNumber(value) {
			return false
		}
		clean += string(value)
	}
	return checkValidNum(clean)
}

func checkValidNum(s string) bool {
	for i := len(s) - 2; i >= 0; i -= 2 {
		num, _ := strconv.Atoi(string(s[i]))
		if num*2 > 9 {
			num = num*2 - 9
		} else {
			num *= 2
		}
		help := strconv.Itoa(num)
		s = s[0:i] + help + s[i+1:len(s)]
	} // s done with double every second digit from the right
	sum := 0 // now sum all digits from s
	for _, value := range s {
		num, _ := strconv.Atoi(string(value))
		sum += num
	}
	if sum%10 != 0 {
		return false
	}
	return true
}

var test27 = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"single digit strings can not be valid",
		"1",
		false,
	},
	{
		"A single zero is invalid",
		"0",
		false,
	},
	{
		"a simple valid SIN that remains valid if reversed",
		"059",
		true,
	},
	{
		"a simple valid SIN that becomes invalid if reversed",
		"59",
		true,
	},
	{
		"a valid Canadian SIN",
		"055 444 285",
		true,
	},
	{
		"invalid Canadian SIN",
		"055 444 286",
		false,
	},
	{
		"invalid credit card",
		"8273 1232 7352 0569",
		false,
	},
	{
		"valid strings with a non-digit included become invalid",
		"055a 444 285",
		false,
	},
	{
		"valid strings with punctuation included become invalid",
		"055-444-285",
		false,
	},
	{
		"valid strings with symbols included become invalid",
		"055£ 444$ 285",
		false,
	},
	{
		"single zero with space is invalid",
		" 0",
		false,
	},
	{
		"more than a single zero is valid",
		"0000 0",
		true,
	},
	{
		"input digit 9 is correctly converted to output digit 9",
		"091",
		true,
	},
}
