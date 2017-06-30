package main

import (
	"fmt"
)

func wordy() {
	fmt.Println("wordy")
}

var test55 = []struct {
	q  string
	a  int
	ok bool
}{
	{"What is 1 plus 1?", 2, true},
	{"What is 53 plus 2?", 55, true},
	{"What is -1 plus -10?", -11, true},
	{"What is 123 plus 45678?", 45801, true},
	{"What is 4 minus -12?", 16, true},
	{"What is -3 multiplied by 25?", -75, true},
	{"What is 33 divided by -3?", -11, true},
	{"What is 1 plus 1 plus 1?", 3, true},
	{"What is 1 plus 5 minus -2?", 8, true},
	{"What is 20 minus 4 minus 13?", 3, true},
	{"What is 17 minus 6 plus 3?", 14, true},
	{"What is 2 multiplied by -2 multiplied by 3?", -12, true},
	{"What is -3 plus 7 multiplied by -2?", -8, true},
	{"What is -12 divided by 2 divided by -3?", 2, true},
	{"What is 53 cubed?", 0, false},
	{"Who is the president of the United States?", 0, false},
}
