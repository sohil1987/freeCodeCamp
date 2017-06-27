package main

import (
	"fmt"
	"regexp"
)

func telephoneCheck(str string) bool {
	pattern := `^(1\s?)?(\(\d{3}\)|\d{3})[\s\-]?\d{3}[\s\-]?\d{4}$`
	res, _ := regexp.MatchString(pattern, str)
	return res
}

func validateUSPhone() {
	fmt.Println("validateUSPhone")

	fmt.Println(telephoneCheck("1 555-555-5555"))
	fmt.Println(telephoneCheck("1 (555) 555-5555"))
	fmt.Println(telephoneCheck("5555555555"))
	fmt.Println(telephoneCheck("555-555-5555"))
	fmt.Println(telephoneCheck("(555)555-5555"))
	fmt.Println(telephoneCheck("1(555)555-5555"))
	fmt.Println(telephoneCheck("555-5555"))
	fmt.Println(telephoneCheck("5555555"))
	fmt.Println(telephoneCheck("1 555 555 5555"))
	fmt.Println(telephoneCheck("1 456 789 4444"))
	fmt.Println(telephoneCheck("123**&!!asdf#"))
	fmt.Println(telephoneCheck("55555555"))
	fmt.Println(telephoneCheck("(6505552368)"))
	fmt.Println(telephoneCheck("2 (757) 622-7382"))
	fmt.Println(telephoneCheck("0 (757) 622-7382"))
	fmt.Println(telephoneCheck("-1 (757) 622-7382"))
	fmt.Println(telephoneCheck("2 757 622-7382"))
	fmt.Println(telephoneCheck("10 (757) 622-7382"))
	fmt.Println(telephoneCheck("27576227382"))
	fmt.Println(telephoneCheck("(275)76227382"))
	fmt.Println(telephoneCheck("2(757)6227382"))
	fmt.Println(telephoneCheck("2(757)622-7382"))
	fmt.Println(telephoneCheck("555)-555-5555"))
	fmt.Println(telephoneCheck("(555-555-5555"))
	fmt.Println(telephoneCheck("(555)5(55?)-5555"))

}

/*
var test = []struct {
}{
}
*/
