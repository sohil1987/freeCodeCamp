package main

import "fmt"

func series() {
	for _, v := range test16 {
		testSerie(v.n, v.s)
		//fmt.Println(v)
	}
}

func testSerie(n int, s string) {
	//fmt.Println(n, s)
	var result []string
	for i := 0; i < len(s); i++ {
		if i+n < len(s)+1 {
			result = append(result, s[i:i+n])
		}
		//fmt.Println(i + n)
	}
	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println("nil")
	}
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
