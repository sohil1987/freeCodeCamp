package main

import (
	"fmt"
	"strings"
)

func scrabbleScore() {
	for _, test := range test20 {
		value := 0
		word := strings.ToLower(test.input)
		for _, letter := range word {
			value += lv[string(letter)]
		}
		fmt.Printf("%-28s value = %3d ,is OK? %t \n", test.input, value, test.expected == value)
	}
}

type t20 struct {
	input    string
	expected int
}

var test20 = []t20{
	{"a", 1},                // lowercase letter
	{"A", 1},                // uppercase letter
	{"f", 4},                // valuable letter
	{"at", 2},               // short word
	{"zoo", 12},             // short, valuable word
	{"street", 6},           // medium word
	{"quirky", 22},          // medium, valuable word
	{"OxyphenButazone", 41}, // long, mixed-case word
	{"pinata", 8},           // english-like word
	{"", 0},                 // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
}

type lettersValues map[string]int

var lv = lettersValues{
	"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1, "d": 2, "g": 2, "b": 3, "c": 3, "m": 3, "p": 3, "f": 4, "h": 4, "v": 4, "w": 4, "y": 4, "k": 5, "j": 8, "x": 8, "q": 10, "z": 10,
}
