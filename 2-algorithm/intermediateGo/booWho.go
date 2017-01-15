package main

import (
	"fmt"
)

func boo(i interface{}) bool {
	fmt.Println("---->>>>>  ", i)
	_, ok := i.(bool)
	return ok
}

func booWho() {
	fmt.Println(boo(true))
	fmt.Println(boo(false))
	fmt.Println(boo([3]int{1, 2, 3}))
	fmt.Println(boo(1))
	fmt.Println(boo("a"))
	fmt.Println(boo("true"))
	fmt.Println(boo("false"))
}
