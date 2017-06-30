package main

import (
	"fmt"
	"reflect"
	"strings"
)

func etl() {
	for _, v := range test19 {
		res := make(map[string]int)
		for value, letters := range v.input {
			for _, letter := range letters {
				res[strings.ToLower(letter)] = value
			}
		}
		fmt.Println(res, v.output)
		eq := reflect.DeepEqual(res, v.output)
		fmt.Printf("is OK? %t\n", eq)
	}
}

type given map[int][]string
type expectation map[string]int

type t19 struct {
	input  given
	output expectation
}

var test19 = []t19{
	{
		given{1: {"A"}},
		expectation{"a": 1},
	},
	{
		given{1: {"A", "E", "I", "O", "U"}},
		expectation{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1},
	},
	{
		given{1: {"A", "E"}, 2: {"D", "G"}},
		expectation{"a": 1, "e": 1, "d": 2, "g": 2},
	},
	{
		given{
			1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
			2:  {"D", "G"},
			3:  {"B", "C", "M", "P"},
			4:  {"F", "H", "V", "W", "Y"},
			5:  {"K"},
			8:  {"J", "X"},
			10: {"Q", "Z"},
		},
		expectation{
			"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
			"d": 2, "g": 2,
			"b": 3, "c": 3, "m": 3, "p": 3,
			"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
			"k": 5,
			"j": 8, "x": 8,
			"q": 10, "z": 10,
		},
	},
}

/*
var test19 = []struct {
	input  given
	output expectation
}{
	{
		given{1: {"A"}},
		expectation{"a": 1},
	},
	{
		given{1: {"A", "E", "I", "O", "U"}},
		expectation{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1},
	},
	{
		given{
			1: {"A", "E"},
			2: {"D", "G"},
		},
		expectation{
			"a": 1,
			"e": 1,
			"d": 2,
			"g": 2,
		},
	},
	{
		given{
			1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
			2:  {"D", "G"},
			3:  {"B", "C", "M", "P"},
			4:  {"F", "H", "V", "W", "Y"},
			5:  {"K"},
			8:  {"J", "X"},
			10: {"Q", "Z"},
		},
		expectation{
			"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
			"d": 2, "g": 2,
			"b": 3, "c": 3, "m": 3, "p": 3,
			"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
			"k": 5,
			"j": 8, "x": 8,
			"q": 10, "z": 10,
		},
	},
}
*/
