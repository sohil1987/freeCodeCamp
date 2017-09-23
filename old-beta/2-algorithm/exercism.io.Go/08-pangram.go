package main

import "fmt"
import "strings"

func pangram() {
	for _, v := range test08 {
		letters := "abcdefghijklmnopqrstuvwxyz"
		v.input = strings.ToLower(v.input)
		pangram := true
		for _, t := range v.input {
			if strings.Contains(letters, string(t)) {
				letters = strings.Replace(letters, string(t), "", 1)
			}
		}
		if letters != "" {
			pangram = false
		}
		fmt.Printf("Solution = %-20t , is OK ? %t\n", pangram, v.expected == pangram)
	}
}

type t08 struct {
	input         string
	expected      bool
	failureReason string
}

var test08 = []t08{
	{"", false, "sentence empty"},
	{"The quick brown fox jumps over the lazy dog", true, ""},
	{"a quick movement of the enemy will jeopardize five gunboats", false, "missing character 'x'"},
	{"the quick brown fish jumps over the lazy dog", false, "another missing character 'x'"},
	{"the 1 quick brown fox jumps over the 2 lazy dogs", true, ""},
	{"7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog", false, "missing letters replaced by numbers"},
	{"\"Five quacking Zephyrs jolt my wax bed.\"", true, ""},
	{"Victor jagt zwölf Boxkämpfer quer über den großen Sylter Deich.", true, ""},
	{"Широкая электрификация южных губерний даст мощный толчок подъёму сельского хозяйства.", false, "Panagram in alphabet other than ASCII"},
}
