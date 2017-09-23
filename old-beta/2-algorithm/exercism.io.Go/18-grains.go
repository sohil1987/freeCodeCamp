package main

import "fmt"
import "math"

func grains() {
	for _, v := range test18 {
		var res uint64
		if v.input > 0 && v.input < 65 {
			res = getGrains18(v.input)
		}
		fmt.Printf("Square = %-3d with %-20d grains, is OK? %t\n", v.input, res, res == v.expectedVal)

	}
}

func getGrains18(box int) uint64 {
	var res uint64
	res = uint64(math.Pow(2, float64(box-1)))
	return res
}

var test18 = []struct {
	input       int
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
	{-1, 0, true},
}
