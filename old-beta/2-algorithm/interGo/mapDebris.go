package main

import "fmt"
import "math"

func mapDebris() {
	var total = make([]sol, len(test07))
	for i, v := range test07 {
		var s sol
		for k, t := range v.data {
			fmt.Sprintln(i, v, k, t)
			var aux response
			aux.name = t.name
			aux.orbitalPeriod = getOrbitalPeriod(t.avgAlt)
			s = append(s, aux)
		}
		total[i] = s
	}
	for i, v := range total {
		fmt.Println(i, v)
	}
}

func getOrbitalPeriod(alt float64) int {
	var gm = 398600.4418
	var radius = 6367.4447
	h := 2 * math.Pi * math.Sqrt((math.Pow(alt+radius, 3) / gm))
	return roundFloat64(h)
}

func roundFloat64(num float64) int {
	if num < 0 {
		return int(num - 0.5)
	}
	return int(num + 0.5)
}

type response struct {
	name          string
	orbitalPeriod int
}

type sol []response

type entry struct {
	name   string
	avgAlt float64
}

var test07 = []struct {
	data []entry
}{
	{
		data: []entry{
			{name: "sputnik", avgAlt: 35873.5553},
		},
	},
	{
		data: []entry{
			{name: "iss", avgAlt: 413.6},
			{name: "hubble", avgAlt: 556.7},
			{name: "moon", avgAlt: 378632.553},
		},
	},
}
