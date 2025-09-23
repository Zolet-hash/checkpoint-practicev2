package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
	}

	args := os.Args[1] + os.Args[2]
	seen := make(map[rune]bool)

	for _, r := range args {
		if !seen[r] {
			seen[r] = true
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
