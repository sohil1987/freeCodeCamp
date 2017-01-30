package main

import (
	"fmt"
	"strings"
)

func pangram() {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i, v := range test08 {
		words := strings.Split(strings.ToLower(v.input), "")
		pangramTest := letters
		for _, let := range words {
			pangramTest = strings.Replace(pangramTest, let, "", -1)
			//fmt.Println("Deleted", let, ", len = ", len(pangramTest))
		}
		if len(pangramTest) == 0 {
			fmt.Println(i, "true")
		} else {
			fmt.Println(i, "false")
		}
	}
}

var test08 = []struct {
	input         string
	expected      bool
	failureReason string
}{
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
