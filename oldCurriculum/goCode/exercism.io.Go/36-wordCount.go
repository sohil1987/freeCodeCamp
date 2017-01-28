package main

import (
	"fmt"
	"strings"
)

func wordCount() {
	for _, v := range test36 {
		words := getWords(v.input)
		count := countWords(words)
		fmt.Println(count)
	}
}

func countWords(words []string) frequency {
	f := make(frequency)
	for _, v := range words {
		//fmt.Println(i, v)
		_, ok := f[strings.ToLower(v)]
		if ok {
			f[strings.ToLower(v)]++
		} else if !ok {
			f[strings.ToLower(v)] = 1
		}
	}
	return f
}

func getWords(str string) []string {
	var res []string
	str += " " // to catch last word
	aux := ""
	for _, v := range str {
		if isValidChar(string(v)) {
			aux += string(v)
		} else if len(aux) > 0 {
			if string(aux[0]) == "'" && string(aux[len(aux)-1]) == "'" {
				aux = aux[1 : len(aux)-1]
			}
			res = append(res, aux)
			aux = ""
		}
	}
	return res
}

func isValidChar(str string) bool {
	letters := "abcdefghijklmnopqrstuvwxyz123456789'ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Contains(letters, str)
}

type frequency map[string]int

var test36 = []struct {
	description string
	input       string
	output      frequency
}{
	{
		"count one word",
		"word",
		frequency{"word": 1},
	},
	{
		"count one of each word",
		"one of each",
		frequency{"each": 1, "of": 1, "one": 1},
	},
	{
		"multiple occurrences of a word",
		"one fish two fish red fish blue fish",
		frequency{"blue": 1, "fish": 4, "one": 1, "red": 1, "two": 1},
	},
	{
		"handles cramped lists",
		"one,two,three",
		frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"handles expanded lists",
		"one,\ntwo,\nthree",
		frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"ignore punctuation",
		"car: carpet as java: javascript!!&@$%^&",
		frequency{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		"include numbers",
		"testing, 1, 2 testing",
		frequency{"1": 1, "2": 1, "testing": 2},
	},
	{
		"normalize case",
		"go Go GO Stop stop",
		frequency{"go": 3, "stop": 2},
	},
	{
		"with apostrophes",
		"First: don't laugh. Then: don't cry.",
		frequency{"cry": 1, "don't": 2, "first": 1, "laugh": 1, "then": 1},
	},
	{
		"with quotations",
		"Joe can't tell between 'large' and large.",
		frequency{"and": 1, "between": 1, "can't": 1, "joe": 1, "large": 2, "tell": 1},
	},
}
