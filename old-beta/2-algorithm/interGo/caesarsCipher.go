package main

import "fmt"

func caesarsCipher() {
	var sol = make([]string, 0) //, len(test02))
	for _, v := range test02 {
		fmt.Println(v)
		sol = append(sol, decipher(v.data))
	}
	for i, v := range sol {
		fmt.Println(i, v)
	}
}

func decipher(s string) string {
	var word = ""
	for _, v := range s {
		if v >= 65 && v < 78 {
			word += string(v + 13)
		} else if v >= 78 && v <= 90 {
			word += string(v - 13)
		} else {
			word += string(v)
		}
	}
	return word
}

var test02 = []struct {
	data string
}{
	{"SERR PBQR PNZC"},
	{"SERR CVMMN!"},
	{"SERR YBIR?"},
	{"GUR DHVPX OEBJA SBK WHZCF BIRE GUR YNML QBT."},
}
