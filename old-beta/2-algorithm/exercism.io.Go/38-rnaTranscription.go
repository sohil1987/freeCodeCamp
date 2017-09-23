package main

import (
	"fmt"
	"reflect"
)

func rnaTranscription() {
	for _, t := range test38 {
		res := getRnaTranscription(t.input)
		fmt.Printf("is OK? %5t  %v\n", reflect.DeepEqual(t.expected, res), t.expected)
	}
}

func getRnaTranscription(s string) string {
	convert := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	var res = ""
	for _, letter := range s {
		res += convert[string(letter)]
	}
	return res
}

var test38 = []struct {
	input    string
	expected string
}{
	// rna complement of cytosine is guanine
	{"C", "G"},

	// rna complement of guanine is cytosine
	{"G", "C"},

	// rna complement of thymine is adenine
	{"T", "A"},

	// rna complement of adenine is uracil
	{"A", "U"},

	// rna complement
	{"ACGTGGTCTTAA", "UGCACCAGAAUU"},
}
