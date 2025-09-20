package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	inword := false   //we're not inside a word
	firstword := true //this is our firstword

	for _, r := range input {
		if r == ' ' || r == '\t' {
			inword = false
		} else {
			if !inword && !firstword {
				z01.PrintRune(' ')
			}
			z01.PrintRune(r)
			inword = true
			firstword = false
		}
	}
	z01.PrintRune('\n')

}
