package main

import (
	"fmt"
	"strings"
)

func pairElement(str string) [][2]string {
	var res [][2]string
	list := strings.Split(str, "")
	var aux [2]string
	for _, v := range list {
		//fmt.Println(i, v)
		aux[0] = v
		aux[1] = pair(v)
		res = append(res, aux)
	}
	//fmt.Println(res)
	return res
}

func pair(str string) string {
	if str == "C" {
		return "G"
	}
	if str == "G" {
		return "C"
	}
	if str == "A" {
		return "T"
	}
	if str == "T" {
		return "A"
	}
	return ""
}

func dnaPairing() {
	fmt.Println(pairElement("ATCGA"))
	fmt.Println(pairElement("TTGAG"))
	fmt.Println(pairElement("CTCTA"))
}
