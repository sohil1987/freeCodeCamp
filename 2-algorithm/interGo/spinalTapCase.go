package main

import (
	"fmt"
	"regexp"
	"strings"
)

func spinalTapCase() {
	var sol = make([]string, len(test14))
	for i, v := range test14 {
		sol[i] = convertToSpinal(v.data)
	}
	for i, v := range sol {
		fmt.Println(i, v)
	}
}

func convertToSpinal(s string) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	r := regexp.MustCompile(" |_") //("/ |_/g")
	s = r.ReplaceAllString(s, `-`)
	var inserts = 0
	for i, v := range s {
		if strings.Contains(strings.ToUpper(letters), string(v)) { // if is upper
			i = i + inserts
			if i > 0 && string(s[i-1]) != "-" {
				s = s[0:i] + "-" + strings.ToLower(string(v)) + s[i+1:]
				inserts++
			} else {
				s = s[0:i] + strings.ToLower(string(v)) + s[i+1:]
			}
		}
	}
	return s
}

var test14 = []struct {
	data string
}{
	{"This Is Spinal Tap"},
	{"thisIsSpinalTap"},
	{"The_Andy_Griffith_Show"},
	{"Teletubbies say Eh-oh"},
	{"AllThe-small Things"},
}
