package main

import (
	"fmt"
	"strings"
)

func missingLetters() {
	var sol = make([]string, 0)
	for _, v := range test08 {
		fmt.Println(v)
		sol = append(sol, getMissingLetter(v.data))
	}
	for i, v := range sol {
		fmt.Printf("%d --> %s \n", i, v)
	}
}

func getMissingLetter(s string) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	var index = strings.Index(letters, string(s[0]))
	for i, v := range s {
		if string(v) != string(letters[index+i]) {
			return string(letters[index+i])
		}
	}

	return "undefined"

}

var test08 = []struct {
	data string
}{
	{"bcdf"},
	{"abce"},
	{"abcdefghjklmno"},
	{"bcd"},
	{"yz"},
}
