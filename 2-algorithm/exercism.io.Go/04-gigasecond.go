package main

import (
	"fmt"
	"time"
)

func gigasecond() {
	for _, v := range test04 {
		start := getParsedStart(v.in)
		want := getParsedStart(v.want)
		diff := want.Sub(start)
		if diff.Seconds() == 1e9 {
			fmt.Println(start, want, "Yes")
		} else {
			fmt.Println(start, want, `Not yet`)
		}
	}
}

func getParsedStart(start string) time.Time {
	layout1 := "2006-01-02" // Layout numbers?
	layout2 := "2006-01-02T15:04:05"
	t, err := time.Parse(layout1, start)
	if err != nil {
		t, err = time.Parse(layout2, start)
	}
	return t
}

var test04 = []struct {
	in   string
	want string
}{
	{
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
	{
		"2015-01-24T22:00:00",
		"2046-10-02T23:46:40",
	},
	{
		"2015-01-24T23:59:59",
		"2046-10-03T01:46:39",
	},
}

/*
type test042 []struct {
	in   string
	want string
}

ac := test042{
	{
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
	{
		"2015-01-24T22:00:00",
		"2046-10-02T23:46:40",
	},
	{
		"2015-01-24T23:59:59",
		"2046-10-03T01:46:39",
	},
}
*/
