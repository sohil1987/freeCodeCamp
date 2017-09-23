package main

import "fmt"

func series() {
	for _, v := range test16 {
		var res = make([]string, 0)
		str := v.s
		size := v.n
		for i := 0; i < len(str); i++ {
			var aux = ""
			if i+size-1 < len(str) {
				aux = str[i : i+size]
				res = append(res, aux)
			}
		}
		//fmt.Println(res, v.out)
		fmt.Printf("Row %#v ,is OK ? %t\n", res, isCorrect16(res, v.out))
	}
}

func isCorrect16(res, expected []string) bool {
	for i, t := range res {
		if t != expected[i] {
			return false
		}
	}
	return true
}

var test16 = []struct {
	n   int
	s   string
	out []string
}{
	{1, "01234",
		[]string{"0", "1", "2", "3", "4"}},
	{1, "92834",
		[]string{"9", "2", "8", "3", "4"}},
	{2, "01234",
		[]string{"01", "12", "23", "34"}},
	{2, "98273463",
		[]string{"98", "82", "27", "73", "34", "46", "63"}},
	{2, "37103",
		[]string{"37", "71", "10", "03"}},
	{3, "01234",
		[]string{"012", "123", "234"}},
	{3, "31001",
		[]string{"310", "100", "001"}},
	{3, "982347",
		[]string{"982", "823", "234", "347"}},
	{4, "01234",
		[]string{"0123", "1234"}},
	{4, "91274",
		[]string{"9127", "1274"}},
	{5, "01234",
		[]string{"01234"}},
	{5, "81228",
		[]string{"81228"}},
}
