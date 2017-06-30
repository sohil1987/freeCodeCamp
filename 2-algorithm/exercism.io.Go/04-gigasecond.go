package main

import (
	"fmt"
	"time"
)

func gigasecond() {
	for _, v := range test04 {
		init := parseStringToTime(v.init)
		expected := parseStringToTime(v.expected)
		end := init.Add(1e9 * 1000 * time.Millisecond)
		if end == expected {
			fmt.Printf("%s --> %s | OK \n", init, expected)
		} else {
			fmt.Printf("%s --> %s | ERROR \n", init, expected)
		}
	}
}

func parseStringToTime(start string) time.Time {
	layout1 := "2006-01-02" // Layout numbers?
	layout2 := "2006-01-02T15:04:05"
	t, err := time.Parse(layout1, start)
	if err != nil {
		t, err = time.Parse(layout2, start)
	}
	return t
}

var test04 = []struct {
	description string
	init        string
	expected    string
}{
	{
		"date only specification of time",
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"second test for date only specification of time",
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"third test for date only specification of time",
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
	{
		"full time specified",
		"2015-01-24T22:00:00",
		"2046-10-02T23:46:40",
	},
	{
		"full time with day roll-over",
		"2015-01-24T23:59:59",
		"2046-10-03T01:46:39",
	},
}
