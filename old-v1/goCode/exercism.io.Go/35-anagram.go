package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram() {
	for _, v := range test35 {
		str := v.subject
		arr := v.candidates
		sol := makeSearch(str, arr)

		fmt.Println(v.subject, " --> ", sol)
	}
}

func makeSearch(str string, arr []string) []string {
	var sol []string
	str = stringToLowerAndSort(str)
	for _, v := range arr {
		if stringToLowerAndSort(v) == str {
			sol = append(sol, v)
		}
	}
	return sol

}

func stringToLowerAndSort(str string) string {
	aux := strings.Split(strings.ToLower(str), "")
	sort.Strings(aux)
	str = strings.Join(aux, "")
	//fmt.Println(str)
	return str
}

var test35 = []struct {
	subject     string
	candidates  []string
	expected    []string
	description string
}{
	{
		subject: "diaper",
		candidates: []string{
			"hello",
			"world",
			"zombies",
			"pants",
		},
		expected:    []string{},
		description: "no matches",
	},
	{
		subject: "ant",
		candidates: []string{
			"tan",
			"stand",
			"at",
		},
		expected:    []string{"tan"},
		description: "simple anagram",
	},
	{
		subject: "listen",
		candidates: []string{
			"enlists",
			"google",
			"inlets",
			"banana",
		},
		expected:    []string{"inlets"},
		description: "another simple anagram",
	},
	{
		subject: "master",
		candidates: []string{
			"stream",
			"pigeon",
			"maters",
		},
		expected:    []string{"maters", "stream"},
		description: "multiple anagrams",
	},
	{
		subject: "allergy",
		candidates: []string{
			"gallery",
			"ballerina",
			"regally",
			"clergy",
			"largely",
			"leading",
		},
		expected:    []string{"gallery", "largely", "regally"},
		description: "multiple anagrams (again)",
	},
	{
		subject: "galea",
		candidates: []string{
			"eagle",
		},
		expected:    []string{},
		description: "does not confuse different duplicates",
	},
	{
		subject: "corn",
		candidates: []string{
			"corn",
			"dark",
			"Corn",
			"rank",
			"CORN",
			"cron",
			"park",
		},
		expected:    []string{"cron"},
		description: "identical word is not anagram",
	},
	{
		subject: "mass",
		candidates: []string{
			"last",
		},
		expected:    []string{},
		description: "eliminate anagrams with same checksum",
	},
	{
		subject: "good",
		candidates: []string{
			"dog",
			"goody",
		},
		expected:    []string{},
		description: "eliminate anagram subsets",
	},
	{
		subject: "Orchestra",
		candidates: []string{
			"cashregiser",
			"carthorse",
			"radishes",
		},
		expected:    []string{"carthorse"},
		description: "subjects are case insensitive",
	},
	{
		subject: "orchestra",
		candidates: []string{
			"cashregiser",
			"Carthorse",
			"radishes",
		},
		expected:    []string{"carthorse"},
		description: "candidates are case insensitive",
	},
}
