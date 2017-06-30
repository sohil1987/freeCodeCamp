package main

import (
	"fmt"
	"strings"
)

func isogram() {
	for _, test := range test24 {
		isIt := isIsogram(test.word)
		fmt.Printf("%-30s is Isogram = %t ,is OK? %t\n", test.word, isIt, isIt == test.expected)
	}
}

func isIsogram(str string) bool {
	exists := ""
	str = strings.Replace(strings.ToLower(str), "-", " ", -1)
	str = strings.Replace(str, " ", "", -1)
	for _, letter := range str {
		if strings.Contains(exists, string(letter)) {
			return false
		}
		exists += string(letter)
	}
	return true
}

var test24 = []struct {
	word     string
	expected bool
}{
	{"duplicates", true},
	{"eleven", false},
	{"subdermatoglyphic", true},
	{"Alphabet", false},
	{"thumbscrew-japingly", true},
	{"Hjelmqvist-Gryb-Zock-Pfund-Wax", true},
	{"Heizölrückstoßabdämpfung", true},
	{"the quick brown fox", false},
	{"Emily Jung Schwartzkopf", true},
	{"éléphant", false},
}
