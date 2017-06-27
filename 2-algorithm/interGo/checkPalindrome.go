package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkPalindrome() {
	var sol = make(map[string](bool))
	for _, v := range test03 {
		sol[v.data] = check(v.data)
	}
	for k, v := range sol {
		fmt.Printf(" %-40s %t \n", k, v)
	}
}

func check(s string) bool {
	var isPalin = true
	s = strings.ToLower(s)
	var clean string
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			clean += string(v)
		}
	}
	var index = 0
	for isPalin == true && index < len(clean) {
		fmt.Printf("%s - %s \n", string(clean[index]), string(clean[len(clean)-1-index]))
		if clean[index] != clean[len(clean)-1-index] {
			isPalin = false
		}
		index++
	}
	return isPalin
}

var test03 = []struct {
	data string
}{
	{"_eye"},
	{"race car"},
	{"not a palindrome"},
	{"A man, a plan, a canal. Panama"},
	{"never odd or even"},
	{"nope"},
	{"almostomla"},
	{"My age is 0, 0 si ega ym."},
	{"1 eye for of 1 eye."},
	{`0_0 (: /-\ :) 0-0`},
	{`five|\_/|four`},
}
