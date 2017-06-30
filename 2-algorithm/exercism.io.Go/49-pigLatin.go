package main

import (
	"fmt"
	"strings"
)

func pigLatin() {
	for _, t := range test49 {
		res := getPigLatin(t.input)
		fmt.Printf("%18s ==> %-18s is OK? %5t\n", t.input, res, t.pig == res)
	}
}

func getPigLatin(s string) string {
	const vowels = "aeiou"
	//if strings.Index(vowels, string(s[0])) != -1 {
	if strings.Contains(vowels, string(s[0])) {
		return s + "ay"
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

var test49 = []struct {
	pig   string
	input string
}{
	{"appleay", "apple"},
	{"earay", "ear"},
	{"igpay", "pig"},
	{"oalakay", "koala"},
	{"airchay", "chair"},
	{"ueenqay", "queen"},
	{"uaresqay", "square"},
	{"erapythay", "therapy"},
	{"ushthray", "thrush"},
	{"oolschay", "school"},
	{"uick fast runqay", "quick fast run"},
	{"ellowyay", "yellow"},
	{"iayttray", "yttria"},
	{"enonxay", "xenon"},
	{"ayxray", "xray"},
}
