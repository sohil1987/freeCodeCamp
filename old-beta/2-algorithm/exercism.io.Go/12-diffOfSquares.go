package main

import "fmt"

func diffOfSquares() {
	for _, v := range test12 {
		var sum, sqr int
		for i := 1; i <= v.n; i++ {
			sum += i
			sqr += i * i
		}
		sum = sum * sum
		fmt.Printf("Solution = %-10d %-10d , is OK ? %t %t\n", sum, sqr, v.sqOfSums == sum, v.sumOfSq == sqr)

	}
}

var test12 = []struct {
	n, sqOfSums, sumOfSq int
}{
	{5, 225, 55},
	{10, 3025, 385},
	{100, 25502500, 338350},
}
