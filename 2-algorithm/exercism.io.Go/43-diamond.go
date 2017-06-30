package main

import (
	"fmt"
	"strings"
)

func diamond() {
	var letters string
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for index, letter := range letters {
		maxRows := 2*index + 1
		sol := drawDiamond(string(letter), index, maxRows, letters)
		printSol(sol)
	}
}

func drawDiamond(str string, index, maxRows int, letters string) []string {
	var sol = make([]string, 0)
	// fill with +
	for row := 0; row < maxRows; row++ {
		text := ""
		for col := 0; col < maxRows; col++ {
			text += "+"
		}
		sol = append(sol, text)
	}
	// put diamond
	center := maxRows / 2

	for row := 0; row < maxRows; row++ {
		if row == 0 {
			sol[row] = replaceAtIndex(sol[row], letters[row], center)
		} else if row > center {
			sol[row] = sol[maxRows-row-1]
		} else {
			sol[row] = replaceAtIndex(sol[row], []byte(letters)[row], center-row)
			sol[row] = replaceAtIndex(sol[row], []byte(letters)[row], center+row)
		}
	}

	fmt.Sprintln(center)
	return sol
}

func replaceAtIndex(input string, rep byte, index int) string {
	var s string
	s = strings.Join([]string{input[:index], string(rep), input[index+1:]}, "")
	return s
}

func printSol(sol []string) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for _, row := range sol {
		fmt.Println(row)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
