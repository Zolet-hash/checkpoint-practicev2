package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[rune]bool)
	existins2 := make(map[rune]bool)

	for _, r := range s2 {
		existins2[r] = true
	}

	for _, ch := range s1 {
		if existins2[ch] && !seen[ch] {
			z01.PrintRune(ch)
			seen[ch] = true
		}
	}
	z01.PrintRune('\n')
}
