package main

import (
	"fmt"
)

func ocrNumbers() {
	for _, t := range test54 {
		fmt.Println(t.in)
		//res := getPrimePos(t.n)
		//fmt.Printf("%10d is in %-6d position, is OK? %5t\n", t.n, res, t.order == res)
	}
}

var test54 = []struct {
	in  string
	out []string
}{
	{`
 _ 
| |
|_|
   `, []string{"0"}},
	{`
   
  |
  |
   `, []string{"1"}},
	{`
 _ 
 _|
|_ 
   `, []string{"2"}},
	{`
 _ 
 _|
 _|
   `, []string{"3"}},
	{`
   
|_|
  |
   `, []string{"4"}},
	{`
 _ 
|_ 
 _|
   `, []string{"5"}},
	{`
 _ 
|_ 
|_|
   `, []string{"6"}},
	{`
 _ 
  |
  |
   `, []string{"7"}},
	{`
 _ 
|_|
|_|
   `, []string{"8"}},
	{`
 _ 
|_|
 _|
   `, []string{"9"}},
	{`
    _ 
  || |
  ||_|
      `, []string{"10"}},
	{`
   
| |
| |
   `, []string{"?"}},
	{`
       _     _        _  _ 
  |  || |  || |  |  || || |
  |  ||_|  ||_|  |  ||_||_|
                           `, []string{"110101100"}},
	{`
       _     _           _ 
  |  || |  || |     || || |
  |  | _|  ||_|  |  ||_||_|
                           `, []string{"11?10?1?0"}},
	{`
    _  _     _  _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|| |
  ||_  _|  | _||_|  ||_| _||_|
                              `, []string{"1234567890"}},
	{`
    _  _ 
  | _| _|
  ||_  _|
         
    _  _ 
|_||_ |_ 
  | _||_|
         
 _  _  _ 
  ||_||_|
  ||_| _|
         `, []string{"123", "456", "789"}},
}
