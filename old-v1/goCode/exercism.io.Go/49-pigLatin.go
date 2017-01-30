package main

import (
	"fmt"
	"strings"
)

func pigLatin() {
	for _, v := range test49 {
		sol := translatePigLatin(v.in)
		fmt.Println(sol)
	}
}

func translatePigLatin(str string) string {
	//vowels := []string{"a", "e", "i", "o", "u"}
	vowels := "aeiou"
	chain := strings.Split(str, "")
	var res string
	if strings.Index(vowels, string(str[0])) != -1 {
		return str + "way"
	}
	for i, v := range chain {
		//fmt.Println(i, v)
		if strings.Index(vowels, v) != -1 {
			res = str[i:] + str[:i] + "ay"
			return res
		}
	}
	return "ERROR"
}

var test49 = []struct {
	pl,
	in string
}{
	{"appleay", "apple"},
	{"earay", "ear"},
	{"igpay", "pig"},
	{"oalakay", "koala"},
	{"airchay", "chair"},
	{"eenquay", "queen"},
	{"aresquay", "square"},
	{"erapythay", "therapy"},
	{"ushthray", "thrush"},
	{"oolschay", "school"},
	{"ickquay astfay unray", "quick fast run"},
	{"ellowyay", "yellow"},
	{"yttriaay", "yttria"},
	{"enonxay", "xenon"},
	{"xrayay", "xray"},
}
