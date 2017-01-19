package main

import "fmt"

func hamming() {
	for i, v := range test05 {
		doTest(v.s1, v.s2, i, v.want) //.s1, v.s2, v.want)
	}
}

func doTest(s1, s2 string, i, w int) {
	dist := 0
	if len(s1) != len(s2) {
		fmt.Println(i, "-1")
		return
	} else if s1 == "" && s2 == "" {
		fmt.Println(i, "0")
		return
	}
	for index := 0; index < len(s1); index++ {
		if s1[index] != s2[index] {
			dist++
		}
	}
	fmt.Println(i, dist, w)
}

var test05 = []struct {
	s1   string
	s2   string
	want int
}{
	{ // identical strands
		"A",
		"A",
		0,
	},
	{ // long identical strands
		"GGACTGA",
		"GGACTGA",
		0,
	},
	{ // complete distance in single nucleotide strands
		"A",
		"G",
		1,
	},
	{ // complete distance in small strands
		"AG",
		"CT",
		2,
	},
	{ // small distance in small strands
		"AT",
		"CT",
		1,
	},
	{ // small distance
		"GGACG",
		"GGTCG",
		1,
	},
	{ // small distance in long strands
		"ACCAGGG",
		"ACTATGG",
		2,
	},
	{ // non-unique character in first strand
		"AGA",
		"AGG",
		1,
	},
	{ // non-unique character in second strand
		"AGG",
		"AGA",
		1,
	},
	{ // large distance
		"GATACA",
		"GCATAA",
		4,
	},
	{ // large distance in off-by-one strand
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
	},
	{ // empty strands
		"",
		"",
		0,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		-1,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		-1,
	},
}
