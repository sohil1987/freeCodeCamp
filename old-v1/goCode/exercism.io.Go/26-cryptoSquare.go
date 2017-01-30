package main

import (
	"fmt"
	"math"
	"strings"
)

type solution struct {
	plain   string
	rawCode string
	rows    int
	cols    int
}

func cryptoSquare() {
	for _, test := range test26 {
		var s solution
		s.plain = test.plain
		s.plain = sanitizeText(s.plain)
		s.cols, s.rows = calculateRectangle26(s.plain)
		aux := prepareArray(&s)
		calculateRawCode26(&s, aux)
		outputEncodedText(s)
		//fmt.Println(s.rawCode)

	}
}

func outputEncodedText(s solution) {
	//fmt.Println(s.rawCode)
	l := len(s.rawCode)
	_, r := calculateRectangle26(s.rawCode)
	if r > 0 {
		//resto := l % r
		fmt.Println("OUTPUT --> ", l, r, "RESTO", l%r)
	}
	var solution string
	//for i, v := range s.rawCode {
	//fmt.Println(i, string(v))
	//fmt.Println("aQUI", i, " / ", r, " = ", i%r, string(v))
	//if i > 0 && i%r == 0 {
	//fmt.Println("aQUI", i, string(v))
	//}
	fmt.Sprintln(r, solution)
	/*if len(s.rawCode) == 0 {
		fmt.Println(s.plain, " == ", "nil")
	} else {
		fmt.Printf("*******************************************\n")
		fmt.Printf("%s == %s\n", s.plain, s.rawCode)
		fmt.Printf("*******************************************\n")
	}*/
}

func calculateRawCode26(s *solution, aux [][]string) {
	// https://rosettacode.org/wiki/Matrix_transposition#Go
	//fmt.Println(len(s.plain), s, aux)
	s.rawCode = ""
	// transpose ,create container
	columns := make([][]string, len(aux[0]))
	for x := range columns {
		columns[x] = make([]string, len(aux))
	}
	// fill columns
	for y, v := range aux {
		for x, t := range v {
			columns[x][y] = t
		}
	}
	//fmt.Println(columns)
	// fill rawcode
	for _, v := range columns {
		for _, t := range v {
			s.rawCode += t
		}
	}
	//fmt.Println(columns)
	//fmt.Println("RAWCODE ----> ", s.rawCode, len(s.rawCode))
}

func prepareArray(s *solution) [][]string {
	aux := strings.Split(s.plain, "")
	var aux2 [][]string
	//fmt.Println(aux)
	unitLen := s.cols // + 1
	i := 0
	for i = 0; i < s.rows-1; i++ {
		aux2 = append(aux2, aux[i*unitLen:(i+1)*unitLen])
	}
	aux2 = append(aux2, aux[i*unitLen:len(aux)])
	//fmt.Println(aux2)
	return aux2
}

func calculateRectangle26(plain string) (int, int) {
	found := false
	aux := 1
	for !found && aux < 10000 {
		if math.Pow(float64(aux), 2) >= float64(len(plain)) {
			found = true
		} else {
			aux++
		}
	}
	cols := aux
	rows := aux - 1
	if cols*rows < len(plain) {
		rows++
	}
	//fmt.Println(len(plain), rows, cols)
	return cols, rows
}

func sanitizeText(plain string) string {
	var res string
	valid := "abcdefghijklmnopqrstuvwxyz"
	plain = strings.ToLower(plain)
	for i := 0; i < len(plain); i++ {
		letter := string(plain[i])
		if strings.Contains(valid, letter) {
			res += letter
		}
	}
	return res
}

var test26 = []struct {
	plain  string // plain text
	cipher string // cipher text
}{
	{
		"Have a nice day. Feed the dog & chill out!",
		"hifei acedl veeol eddgo aatcu nyhht",
	},
	{
		"s#$%^&plunk",
		"su pn lk",
	},
	{
		"1, 2, 3 GO!",
		"1g 2o 3",
	},
	{
		"1234",
		"13 24",
	},
	{
		"123456789",
		"147 258 369",
	},
	{
		"123456789abc",
		"159 26a 37b 48c",
	},
	{
		"Never vex thine heart with idle woes",
		"neewl exhie vtetw ehaho ririe vntds",
	},
	{
		"ZOMG ZOMBIES!!!",
		"zzi ooe mms gb",
	},
	{
		"Time is an illusion. Lunchtime doubly so.",
		"tasney inicds miohoo elntu illib suuml",
	},
	{
		"We all know interspecies romance is weird.",
		"wneiaw eorene awssci liprer lneoid ktcms",
	},
	{
		"Madness, and then illumination.",
		"msemo aanin dnin ndla etlt shui",
	},
	{
		"Vampires are people too!",
		"vrel aepe mset paoo irpo",
	},
	{
		"",
		"",
	},
	{
		"1",
		"1",
	},
	{
		"12",
		"1 2",
	},
	{
		"12 3",
		"13 2",
	},
	{
		"12345678",
		"147 258 36",
	},
	{
		"123456789a",
		"159 26a 37 48",
	},
	{
		"If man was meant to stay on the ground god would have given us roots",
		"imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau",
	},
}
