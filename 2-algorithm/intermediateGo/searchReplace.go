package main

import (
	"fmt"
	"strings"
)

func myReplace(str, before, after string) string {
	var result string
	if string(before[0]) == strings.ToUpper(string(before[0])) {
		//fmt.Println(before[0], string(before[0]))
		after = strings.ToUpper(string(after[0])) + after[1:len(after)]
	}
	result = strings.Replace(str, before, after, 1)
	return result
}

func searchReplace() {
	fmt.Println(myReplace("Let us go to the store", "store", "mall"))
	fmt.Println(myReplace("He is Sleeping on the couch", "Sleeping", "sitting"))
	fmt.Println(myReplace("This has a spellngi error", "spellngi", "spelling"))
	fmt.Println(myReplace("His name is Tom", "Tom", "john"))
	fmt.Println(myReplace("Let us get back to more Coding", "Coding", "algorithms"))
}
