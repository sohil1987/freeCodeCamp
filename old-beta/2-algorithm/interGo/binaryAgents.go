package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryAgents() {
	var sol = make([]string, len(test01))
	for i, v := range test01 {
		letters := strings.Split(v.data, " ")
		word := ""
		for _, letter := range letters {
			next, _ := strconv.ParseInt(letter, 2, 8)
			word = word + string(next)
		}
		sol[i] = word
		word = ""
	}
	for i, v := range sol {
		fmt.Println(i, v)
	}
}

var test01 = []struct {
	data string
}{
	{"01000001 01110010 01100101 01101110 00100111 01110100 00100000 01100010 01101111 01101110 01100110 01101001 01110010 01100101 01110011 00100000 01100110 01110101 01101110 00100001 00111111"},
	{"01001001 00100000 01101100 01101111 01110110 01100101 00100000 01000110 01110010 01100101 01100101 01000011 01101111 01100100 01100101 01000011 01100001 01101101 01110000 00100001"},
}
