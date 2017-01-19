package main

import (
	"fmt"
	"strings"
)

func scrabbleScore() {
	for _, str := range test20 {
		points := 0
		for _, char := range str.input {
			letter := strings.ToLower((string(char)))
			points += values[letter]
		}
		fmt.Println(str.input, points)
	}
}

var values = map[string]int{
	"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
	"d": 2, "g": 2,
	"b": 3, "c": 3, "m": 3, "p": 3,
	"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
	"k": 5,
	"j": 8, "x": 8,
	"q": 10, "z": 10,
}

var test20 = []struct {
	input    string
	expected int
}{
	{"", 0},
	{" \t\n", 0},
	{"a", 1},
	{"f", 4},
	{"street", 6},
	{"quirky", 22},
	{"OXYPHENBUTAZONE", 41},
	{"alacrity", 13},
}
