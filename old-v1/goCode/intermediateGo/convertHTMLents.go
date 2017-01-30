package main

import (
	"fmt"
	"strings"
)

func convertHTML(str string) string {
	var res string
	letters := strings.Split(str, "")
	for _, v := range letters {
		v = convert2(v)
		res += v
	}
	return res
}

func convert2(str string) string {
	if str == "<" {
		return "&lt;"
	}
	if str == ">" {
		return "&gt;"
	}
	if str == "&" {
		return "&amp;"
	}
	if str == "\"" {
		return "&quot;"
	}
	if str == "'" {
		return "&apos;"
	}
	return str
}

func convertHTMLents() {
	fmt.Println(convertHTML("Dolce & Gabbana"))
	fmt.Println(convertHTML("Hamburgers < Pizza < Tacos"))
	fmt.Println(convertHTML("Sixty > twelve"))
	fmt.Println(convertHTML("Stuff in \"quotation marks\""))
	fmt.Println(convertHTML("Shindler's List"))
	fmt.Println(convertHTML("<>"))
	fmt.Println(convertHTML("abc"))
}
