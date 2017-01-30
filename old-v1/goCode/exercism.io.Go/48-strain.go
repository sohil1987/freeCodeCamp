package main

import (
	"fmt"
)

func strain() {
	fmt.Println("strain")
}

/*
(Ints) Keep(func(int) bool) Ints
(Ints) Discard(func(int) bool) Ints
(Lists) Keep(func([]int) bool) Lists
(Strings) Keep(func(string) bool) Strings

type Ints []int
type Lists [][]int
type Strings []string

var tests48 = []struct {
	pred func(int) bool
	list Ints
	want Ints
}{
	{lt10,
		nil,
		nil},
	{lt10,
		Ints{1, 2, 3},
		Ints{1, 2, 3}},
	{odd,
		Ints{1, 2, 3},
		Ints{1, 3}},
	{even,
		Ints{1, 2, 3, 4, 5},
		Ints{2, 4}},
}
*/
