package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	u "utils"
)

func luhn() {
	for _, v := range test27 {
		str := strings.Replace(v.input, " ", "", -1)
		if len(str) > 1 && isAllDigits(str) {
			reverseStr(&str)
			isValid := doubleDigitsAndSum(str)
			fmt.Println(v.input, isValid)
		} else {
			fmt.Println(v.input, "false")
		}
	}
}

func doubleDigitsAndSum(str string) bool {
	var total = 0
	for i, v := range str {
		num, _ := strconv.Atoi(string(v))
		if (i+1)%2 == 0 {
			if num*2 > 9 {
				num = (num * 2) - 9
			} else {
				num = num * 2
			}
			total += num
		} else {
			total += num
		}
	}
	//fmt.Println(str, new, total)
	if total%10 == 0 {
		return true
	}
	return false
}

func reverseStr(str *string) {
	*str = strings.Join(u.ReverseSliceString(strings.Split(*str, "")), "")
}

func isAllDigits(str string) bool {
	for _, v := range str {
		if !unicode.IsNumber(v) {
			//fmt.Println("MECC", v, string(v))
			return false
		}
	}
	return true
}

var test27 = []struct {
	input       string
	description string
	ok          bool
}{
	{"1  ", "single digit strings can not be valid", false},
	{" 0 ", "a single zero is invalid", false},
	{"046 454 286", "valid Canadian SIN", true},
	{"046 454 287", "invalid Canadian SIN", false},
	{"8273 1232 7352 0569", "invalid credit card", false},
	{"827a 1232 7352 0569", "strings that contain non-digits are not valid", false},
}
