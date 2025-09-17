package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	result := []rune{}

	for _, r := range s {
		var count int

		if r >= 'a' && r <= 'z' {
			count = int(r - 'a' + 1)

		} else if r >= 'A' && r <= 'Z' {
			count = int(r - 'A' + 1)

		} else {
			count = 1
		}

		for i := 0; i < count; i++ {
			result = append(result, r)
		}
	}
	return string(result)
}
