package main

import (
	"fmt"
	"strings"
	"unicode"
)

func cryptoSquare() {
	for _, test := range test26 {
		done := cipherText(test.plain)
		fmt.Printf("%-70s \n %-70s ,is OK? %t\n", test.plain, done, done == test.cipher)
	}
}

func cipherText(plain string) string {
	if plain == "" {
		return ""
	}
	plain = sanitizedString(plain)
	fmt.Println(plain)
	cipher := ""
	return cipher
}

func sanitizedString(s string) string {
	s = strings.ToLower(s)
	clean := ""
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			clean += string(letter)
		}
	}
	return clean
}

var test26 = []struct {
	plain  string // plain text
	cipher string // cipher text
}{
	/*	{
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
			"ZOMG! ZOMBIES!!!",
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
		},*/
	{
		"Have a nice day. Feed the dog & chill out!",
		"hifei acedl veeol eddgo aatcu nyhht",
	},
}
