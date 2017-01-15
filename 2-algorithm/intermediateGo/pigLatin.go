package main

import (
	"fmt"
	"strings"
)

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

func pigLatin() {
	fmt.Println(translatePigLatin("california"))
	fmt.Println(translatePigLatin("paragraphs"))
	fmt.Println(translatePigLatin("glove"))
	fmt.Println(translatePigLatin("algorithm"))
	fmt.Println(translatePigLatin("eight"))
}

/*
// only takes first consonant, non consonant cluster
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
		if strings.Index(vowels, v) == -1 {
			res += str[i+1:] + v + "ay"
			return res
		}
		res += v
	}
	return "PigLatin --> " + res
}

*/
