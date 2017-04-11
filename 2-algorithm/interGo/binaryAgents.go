package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func convert(str string) string {
	var message string
	letters := strings.Split(str, " ")

	for _, v := range letters {
		l, err := strconv.ParseInt(v, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		message += fmt.Sprintf("%c", l)
	}

	return message
}

func binaryAgents() {
	fmt.Println(convert("01000001 01110010 01100101 01101110 00100111 01110100 00100000 01100010 01101111 01101110 01100110 01101001 01110010 01100101 01110011 00100000 01100110 01110101 01101110 00100001 00111111"))
	fmt.Println(convert("01001001 00100000 01101100 01101111 01110110 01100101 00100000 01000110 01110010 01100101 01100101 01000011 01101111 01100100 01100101 01000011 01100001 01101101 01110000 00100001"))
}
