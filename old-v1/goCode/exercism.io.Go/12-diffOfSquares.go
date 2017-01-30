package main

import (
	"fmt"
)

func diffOfSquares() {
	for _, v := range test12 {
		var sumSq int
		var sqrSum int
		for i := 1; i <= v.n; i++ {
			sumSq += i * i
			sqrSum += i
		}
		sqrSum = sqrSum * sqrSum
		fmt.Println(v.n, sqrSum, sumSq)
	}
}

var test12 = []struct {
	n, sqOfSums, sumOfSq int
}{
	{5, 225, 55},
	{10, 3025, 385},
	{100, 25502500, 338350},
}
