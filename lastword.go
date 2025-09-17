package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	length := len(s)
	i := length - 1 // obtaining the last index

	for i >= 0 && s[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}

	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return s[i+1:end+1] + "\n"
}
