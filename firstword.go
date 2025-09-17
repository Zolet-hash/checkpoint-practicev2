package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	start := 0 //index
	length := len(s)

	for start < length && s[start] == ' ' {
		start++
	}
	if start == length {
		return s
	}

	end := start
	for end < length && s[end] != ' ' {
		end++
	}
	return s[start:end] + "\n"
}
