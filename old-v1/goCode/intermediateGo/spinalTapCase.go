package main

import (
	"fmt"
	"strings"
)

func spinalCase(str string) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	var result string
	//fmt.Println(str)
	for i, v := range str {
		//fmt.Println(i, string(v))
		var may = strings.ToUpper(string(v))
		if strings.Contains(letters, strings.ToLower(string(v))) {
			if string(v) == may && i > 0 && strings.Contains(letters, string(str[i-1])) {
				result += "-" + string(v)
			} else {
				result += string(v)
			}
		} else {
			result += "-"
		}
	}
	return strings.ToLower(result)
}

func spinalTapCase() {
	fmt.Println(spinalCase("This Is Spinal Tap"))
	fmt.Println(spinalCase("thisIsSpinalTap"))
	fmt.Println(spinalCase("The_Andy_Griffith_Show"))
	fmt.Println(spinalCase("Teletubbies say Eh-oh"))
	fmt.Println(spinalCase("AllThe-small Things"))
}
