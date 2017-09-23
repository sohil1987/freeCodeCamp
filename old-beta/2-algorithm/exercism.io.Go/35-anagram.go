package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram() {
	for _, t := range test35 {
		res := getAnagrams(strings.ToLower(t.subject), t.candidates)
		fmt.Printf("%-12s has %d anagrams ,is OK? %t\n", t.subject, len(res), len(res) == len(t.expected))

	}
}
func getAnagrams(s string, options []string) []string {
	res := make([]string, 0)
	for _, option := range options {
		if isAnagram(s, option) && s != strings.ToLower(option) {
			res = append(res, option)
		}
	}
	return res
}

func isAnagram(s, option string) bool {
	if len(s) != len(option) {
		return false
	}
	s1 := strings.Split(s, "")
	option1 := strings.Split(strings.ToLower(option), "")
	sort.Strings(s1)
	sort.Strings(option1)
	for i, letter := range s1 {
		if letter != string(option1[i]) {
			return false
		}
	}
	return true
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
		expected:    []string{"Carthorse"},
		description: "candidates are case insensitive",
	},
}
