package main

import (
	"fmt"
	"math"
	"strings"
)

func queenAttack() {
	for i := range test17 {
		//fmt.Println(i, v)
		check(i)
		fmt.Println(test17[i].attack)
	}
}

func check(i int) string {
	rows := "abcdefgh"
	cols := "12345678"
	t := &test17[i]
	if t.w == t.b || t.w == "" || t.b == "" || len(t.w) != 2 || len(t.b) != 2 {
		t.attack = false
	} else if t.w[0] == t.b[0] {
		t.attack = true
	} else if t.w[1] == t.b[1] {
		t.attack = true
	} else {
		distRows := strings.Index(rows, string(t.w[0])) - strings.Index(rows, string(t.b[0]))
		distCols := strings.Index(cols, string(t.w[1])) - strings.Index(cols, string(t.b[1]))
		if math.Abs(float64(distRows)) == math.Abs(float64(distCols)) {
			t.attack = true
		} else {
			t.attack = false
		}
	}
	return "nil"
}

var test17 = []struct {
	w, b   string
	attack bool
	ok     bool
}{
	{"b4", "b4", false, false},      // same square
	{"a8", "b9", false, false},      // off board
	{"here", "there", false, false}, // invalid
	{"", "", false, false},
	{"b3", "d7", false, true}, // no attack
	{"b4", "b7", true, true},  // same file
	{"e4", "b4", true, true},  // same rank
	{"a1", "f6", true, true},  // common diagonals
	{"a6", "b7", true, true},
	{"d1", "f3", true, true},
	{"f1", "a6", true, true},
}
