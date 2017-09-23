package main

import (
	"fmt"
	"strings"
)

func searchReplace() {
	var sol = make([]string, 0)
	for _, v := range test11 {
		if string(v.before[0]) == strings.ToUpper(string(v.before[0])) {
			v.after = strings.ToUpper(string(v.after[0])) + v.after[1:]
		}
		v.str = strings.Replace(v.str, v.before, v.after, 1)
		sol = append(sol, v.str)
	}

	for _, v := range sol {
		fmt.Println(v)
	}

}

var test11 = []struct {
	str    string
	before string
	after  string
}{
	{"Let us go to the store", "store", "mall"},
	{"He is Sleeping on the couch", "Sleeping", "sitting"},
	{"This has a spellngi error", "spellngi", "spelling"},
	{"His name is Tom", "Tom", "john"},
	{"Let us get back to more Coding", "Coding", "algorithms"},
}
