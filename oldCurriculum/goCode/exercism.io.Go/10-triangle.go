package main

import (
	"fmt"
	"sort"
)

func triangle() {
	for i, v := range test10 {
		var class string
		var sides = []float64{v.a, v.b, v.c}
		sort.Float64s(sides)
		if v.a <= 0 || v.b <= 0 || v.c <= 0 {
			class = "zero or negative"
		} else if sides[0]+sides[1] < sides[2] {
			class = "inequality"
		} else if v.a == v.b && v.a == v.c {
			class = "equilateral"
		} else if v.a != v.b && v.a != v.c && v.b != v.c {
			class = "scalene"
		} else {
			class = "isosceles"
		}
		fmt.Println(i, class)
	}
}

var test10 = []struct {
	a, b, c float64
}{
	{2, 2, 2},    // same length
	{10, 10, 10}, // a little bigger
	{3, 4, 4},    // last two sides equal
	{4, 3, 4},    // first and last sides equal
	{4, 4, 3},    // first two sides equal
	{10, 10, 2},  // again
	{2, 4, 2},    // a "triangle" that is just a line is still OK
	{3, 4, 5},    // no sides equal
	{10, 11, 12}, // again
	{5, 4, 2},    // descending order
	{.4, .6, .3}, // small sides
	{1, 4, 3},    // a "triangle" that is just a line is still OK
	{5, 4, 6},    // 2a == b+c looks like equilateral, but isn't always.
	{6, 4, 5},    // 2a == b+c looks like equilateral, but isn't always.
	{0, 0, 0},    // zero length
	{3, 4, -5},   // negative length
	{1, 1, 3},    // fails triangle inequality
	{2, 5, 2},    // another
	{7, 3, 2},    // another
}
