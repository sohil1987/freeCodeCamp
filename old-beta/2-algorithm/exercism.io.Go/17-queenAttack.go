package main

import "fmt"
import "strconv"
import "strings"
import "math"

func queenAttack() {
	for i, v := range test17 {
		res := canQueenAttack(i)
		fmt.Printf("Row %-8s,%-8s can attack ? %-5t ,is OK ? %t\n", v.w, v.b, res, res == v.attack)
	}
}

func canQueenAttack(i int) bool {
	rows := "abcdefgh"
	var t = test17[i]
	if len(t.w) != 2 || len(t.b) != 2 {
		return false
	}
	x1 := string(t.w[0])
	x2, err := strconv.Atoi(string(t.w[1]))
	if err != nil {
		return false
	}
	y1 := string(t.b[0])
	y2, err := strconv.Atoi(string(t.b[1]))
	if err != nil {
		return false
	}
	if !strings.Contains(rows, x1) || !strings.Contains(rows, y1) {
		return false
	}
	if x1 == y1 && x2 == y2 {
		return false
	}
	if x2 < 1 || x2 > 8 || y2 < 1 || y2 > 8 {
		return false
	}
	if x1 == y1 || x2 == y2 { // same row or col
		return true
	}
	distY := int(math.Abs(float64(x2) - float64(y2)))
	a := strings.Index(rows, x1)
	b := strings.Index(rows, y1)
	distX := int(math.Abs(float64(a) - float64(b)))
	if distX == distY {
		return true
	}
	return false
}

var test17 = []struct {
	w, b   string
	attack bool
	ok     bool
}{
	{"b4", "b4", false, false}, // same square
	{"a8", "b9", false, false}, // off board
	{"a0", "b1", false, false},
	{"g3", "i5", false, false},
	//{"here", "there", false, false}, // invalid
	//{"", "", false, false},
	//{"z4", "a4", false, false},
	//{"44", "aa", false, false},
	{"b3", "d7", false, true}, // no attack
	{"a1", "f8", false, true},
	{"b4", "b7", true, true}, // same file
	{"e4", "b4", true, true}, // same rank
	{"a1", "f6", true, true}, // common diagonals
	{"a6", "b7", true, true},
	{"d1", "f3", true, true},
	{"f1", "a6", true, true},
	{"a1", "h8", true, true},
	{"a8", "h1", true, true},
}
