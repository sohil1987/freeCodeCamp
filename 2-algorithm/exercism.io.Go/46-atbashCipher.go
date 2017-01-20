package main

import "fmt"
import "strings"

func atbashCipher() {
	for _, v := range test46 {
		sol := decipherAtbash(v.expected)
		fmt.Println(sol)
	}
}

func decipherAtbash(str string) string {
	var sol string
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	for _, v := range str {
		pos := strings.Index(plain, string(v))
		if pos != -1 {
			sol += string(cipher[pos])
		} else if pos == -1 {
			sol += string(v)
		}
	}
	return sol
}

var test46 = []struct {
	expected string
	s        string
}{
	{"ml", "no"},
	{"ml", "no"},
	{"bvh", "yes"},
	{"lnt", "OMG"},
	{"lnt", "O M G"},
	{"nrmwy oldrm tob", "mindblowingly"},
	{"gvhgr mt123 gvhgr mt", "Testing, 1 2 3, testing."},
	{"gifgs rhurx grlm", "Truth is fiction."},
	{"gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt", "The quick brown fox jumps over the lazy dog."},
}
