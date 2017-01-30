package main

import "fmt"

func rnaTranscription() {
	for _, v := range test38 {
		rna := transcriptRna(v.input)
		fmt.Println(rna)
	}
}

func transcriptRna(str string) string {
	convert := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	var rna string
	for _, v := range str {
		rna += convert[string(v)]
	}
	return rna
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
