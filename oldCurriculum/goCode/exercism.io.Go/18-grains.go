package main

import "fmt"

func grains() {
	squares := uint(65)
	total := make([]uint, squares)
	partial := make([]uint, squares)
	calculate(total, partial, squares)
	//fmt.Println(total, "\n", partial)

	for _, v := range test18 {
		if v.input > 0 && v.input < 65 {
			fmt.Println(v.input, partial[v.input-1], total[v.input-1])
		}
	}
}

func calculate(t, p []uint, squares uint) {
	t[0], p[0] = 1, 1
	for i := uint(1); i < squares; i++ {
		p[i] = 2 * p[i-1]
		t[i] = t[i-1] + p[i]
	}
}

var test18 = []struct {
	input       uint
	expectedVal uint64
	expectError bool
}{
	{1, 1, false},
	{2, 2, false},
	{3, 4, false},
	{4, 8, false},
	{16, 32768, false},
	{32, 2147483648, false},
	{64, 9223372036854775808, false},
	{65, 0, true},
	{0, 0, true},
	{100, 0, true},
}
