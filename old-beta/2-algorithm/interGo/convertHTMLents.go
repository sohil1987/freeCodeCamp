package main

import "fmt"

func convertHTMLents() {
	var sol = make([]string, 0)
	for _, v := range test04 {
		var clean = ""
		for _, l := range v.data {
			if string(l) == "<" {
				clean += "&lt;"
			} else if string(l) == ">" {
				clean += "&gt;"
			} else if string(l) == "&" {
				clean += "&amp;"
			} else if string(l) == `"` {
				clean += "&quot;"
			} else if string(l) == `'` {
				clean += "&apos;"
			} else {
				clean += string(l)
			}
		}
		sol = append(sol, clean)
	}
	for _, v := range sol {
		fmt.Printf(" %-40s \n", v)
	}
}

var test04 = []struct {
	data string
}{
	{"Dolce & Gabbana"},
	{"Hamburgers < Pizza < Tacos"},
	{"Sixty > twelve"},
	{"Stuff in \"quotation marks\""},
	{"Shindler's List"},
	{"<>"},
	{"abc"},
}
