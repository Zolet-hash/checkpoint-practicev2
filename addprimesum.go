package main

import (
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printDigit(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}

}
