package main

import (
	"fmt"
	"unicode"
)

func phoneNumber() {
	for _, v := range test47 {
		sol := sanitizeText47(v.input)
		isValidPhone := isValidPhone(sol)
		fmt.Println(v.input, " --> ", isValidPhone)
	}
}

func sanitizeText47(str string) string {
	var res string
	for _, v := range str {
		if unicode.IsDigit(v) {
			res += string(v)
		}
	}
	return res
}

func isValidPhone(str string) bool {
	if len(str) < 10 {
		return false
	} else if len(str) == 10 {
		return true
	} else if len(str) == 11 {
		if str[0] == 1 {
			return true
		}
		return false
	} else if len(str) > 11 {
		return false
	}
	return false
}

var test47 = []struct {
	input     string
	expected  string
	expectErr bool
}{
	{"(123) 456-7890", "1234567890", false},
	{"123.456.7890", "1234567890", false},
	{"1234567890", "1234567890", false},
	{"12345678901234567", "", true},
	{"21234567890", "", true},
	{"123456789", "", true},
}
