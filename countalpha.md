countalpha
Instructions

Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.
Expected functions

func CountAlpha(s string) int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

And its output:

$ go run .
10
5
5