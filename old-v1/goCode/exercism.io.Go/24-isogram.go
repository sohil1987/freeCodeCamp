package main

import (
	"fmt"
	"strings"
)

func isogram() {
	for _, v := range test24 {
		fmt.Println(count24(strings.ToLower(v.word)))
	}
}

func count24(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	counter := make(map[string]int)
	for _, letter := range strings.Split(str, "") {
		letter = strings.ToLower(letter)
		_, ok := counter[letter]
		if !ok {
			counter[letter] = 1
		} else {
			//counter[letter]++
			fmt.Println(letter)
			return false
		}
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
