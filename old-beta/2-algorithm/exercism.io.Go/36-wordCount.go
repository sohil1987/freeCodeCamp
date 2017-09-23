package main

import "fmt"
import "reflect"
import "strings"

func wordCount() {
	for _, t := range test36 {
		res := make(map[string]int)
		res = getWordCount(t.input)
		fmt.Printf("is OK? %5t  %v\n", reflect.DeepEqual(t.output, res), t.output)
	}
}

func getWordCount(s string) map[string]int {
	var res = make(map[string]int)
	words := getWords(strings.ToLower(s))
	for _, word := range words {
		_, exists := res[word]
		if !exists {
			res[word] = 1
		} else {
			res[word]++
		}
	}
	fmt.Println(res)
	return res
}

func getWords(s string) []string {
	validLetters := "abcdefghijklmnopqrstuvwxyz123456789'"
	var words []string
	var clean = ""
	for i, letter := range s {
		fmt.Sprintln(i, string(letter))
		if strings.Contains(validLetters, string(letter)) {
			if string(letter) == "'" {
				if strings.Contains(validLetters, string(s[i-1])) && strings.Contains(validLetters, string(s[i+1])) {
					clean += string(letter)
				}
			} else {
				clean += string(letter)
			}
		} else if i != 0 && i != len(s)-1 { // not first not last
			if string(clean[len(clean)-1]) != "+" {
				clean += "+"
			}
		}
	}
	if string(clean[len(clean)-1]) == "+" {
		clean = clean[1 : len(clean)-1]
	}
	if string(clean[0]) == "+" {
		clean = clean[1:len(clean)]
	}
	words = strings.Split(clean, "+")
	return words
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
