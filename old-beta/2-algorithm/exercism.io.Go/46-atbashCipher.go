package main

import (
	"fmt"
	"reflect"
	"strings"
)

func atbashCipher() {
	for _, t := range test46 {
		res := getAtbashCipher(t.s)
		//fmt.Println(`we get --`, res)
		//fmt.Println(`we want--`, t.expected)
		//fmt.Println(len(t.expected), len(res))
		//fmt.Println(`*************************************`)
		fmt.Printf("is OK? %5t  %v\n", reflect.DeepEqual(t.expected, res), t.expected)
	}
}

func getAtbashCipher(s string) string {
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	nums := "0123456789"
	var res = ""
	s = strings.ToLower(s)
	for _, letter := range s {
		if len(res) > 0 && (len(res)+1)%6 == 0 {
			res += " "
		}
		if strings.Contains(nums, string(letter)) {
			res += string(letter)
		} else if strings.Contains(plain, string(letter)) {
			res += string(cipher[strings.Index(plain, string(letter))])
		}
	}
	if string(res[len(res)-1]) == " " {
		res = res[0 : len(res)-1]
	}
	return res
}

type t46 struct {
	s        string
	expected string
}

var test46 = []t46{
	{
		s:        "yes",
		expected: "bvh",
	},
	{
		s:        "no",
		expected: "ml",
	},
	{
		s:        "OMG",
		expected: "lnt",
	},
	{
		s:        "O M G",
		expected: "lnt",
	},
	{
		s:        "mindblowingly",
		expected: "nrmwy oldrm tob",
	},
	{
		s:        "Testing,1 2 3, testing.",
		expected: "gvhgr mt123 gvhgr mt",
	},
	{
		s:        "Truth is fiction.",
		expected: "gifgs rhurx grlm",
	},
	{
		s:        "The quick brown fox jumps over the lazy dog.",
		expected: "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
	{
		s:        "exercism",
		expected: "vcvix rhn",
	},
	{
		s:        "anobstacleisoftenasteppingstone",
		expected: "zmlyh gzxov rhlug vmzhg vkkrm thglm v",
	},
	{
		s:        "testing123testing",
		expected: "gvhgr mt123 gvhgr mt",
	},
	{
		s:        "thequickbrownfoxjumpsoverthelazydog",
		expected: "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
}
