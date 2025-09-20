package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	original := []rune(os.Args[1])
	oldChar := []rune(os.Args[2])
	newChar := []rune(os.Args[3])

	if len(oldChar) != 1 || len(newChar) != 1 {
		return
	}

	for _, r := range original {
		if == oldChar[0] {
			z01.PrintRune(newChar[0])
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
