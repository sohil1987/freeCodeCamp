package main

import (
	"fmt"
	"strings"
)

func fearNotLetter(str string) string {
	var missing = ""
	letters := "abcdefghijklmnopqrstuvwxyz"
	test := strings.Split(letters, "")
	//fmt.Println(len(test))
	begin := strings.Index(letters, string(str[0]))
	//fmt.Println(begin)
	for i := 0; i < len(str); i++ {
		//fmt.Println("Compare ...", string(str[i]), " con ", test[begin+i])
		if string(str[i]) != test[begin+i] {
			missing = string(test[begin+i])
			return "On " + str + " Fails ... " + missing
		}
	}
	return "On " + str + " nobody fails"
}

func missingLetters() {
	fmt.Println(fearNotLetter("bcdf"))
	fmt.Println(fearNotLetter("abce"))
	fmt.Println(fearNotLetter("abcdefghjklmno"))
	fmt.Println(fearNotLetter("bcd"))
	fmt.Println(fearNotLetter("yz"))
}
