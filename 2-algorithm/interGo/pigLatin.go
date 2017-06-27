package main

import "fmt"
import "strings"

func pigLatin() {
	var sol = make([]string, 0)
	for _, v := range test09 {
		sol = append(sol, getPigLatin(v.data))
	}
	for i, v := range sol {
		fmt.Printf("%d --> %s \n", i, v)
	}
}

func getPigLatin(s string) string {
	const vowels = "aeiou"
	//if strings.Index(vowels, string(s[0])) != -1 {
	if strings.Contains(vowels, string(s[0])) {
		return s + "way"
	}
	var aux string
	for i, v := range s {
		if !strings.Contains(vowels, string(v)) {
			aux += string(v)
		} else {
			return s[i:] + aux + "ay"
		}
	}
	return s
}

var test09 = []struct {
	data string
}{
	{"california"},
	{"paragraphs"},
	{"glove"},
	{"algorithm"},
	{"eight"},
}
