package main

import "fmt"

func dnaPairing() {
	var sol = make([][][]string, len(test06))
	for i, v := range test06 {
		for _, t := range v.data {
			var aux = make([]string, 0)
			aux = append(aux, string(t))
			switch string(t) {
			case "A":
				aux = append(aux, "T")
			case "T":
				aux = append(aux, "A")
			case "C":
				aux = append(aux, "G")
			case "G":
				aux = append(aux, "C")
			}
			sol[i] = append(sol[i], aux)
			aux = make([]string, 0)
		}
	}
	for i, v := range sol {
		fmt.Printf("%s ==> %s \n", test06[i].data, v)
	}
}

var test06 = []struct {
	data string
}{
	{"ATCGA"},
	{"TTGAG"},
	{"CTCTA"},
}
