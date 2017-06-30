package main

import (
	"fmt"
)

func hamming() {
	for _, v := range test05 {
		var dist = 0
		if len(v.s1) != len(v.s2) {
			dist = -1
			fmt.Printf("Solution = %-5d , is OK ? %t\n", dist, v.expected == dist)
		} else {
			dist = 0
			for k, t := range v.s1 {
				if byte(t) != v.s2[k] {
					dist++
				}
			}
			fmt.Printf("Solution = %-5d , is OK ? %t\n", dist, v.expected == dist)
		}
	}
}

var test05 = []struct {
	s1       string
	s2       string
	expected int
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
	{ // same nucleotides in different positions
		"TAG",
		"GAT",
		2,
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
