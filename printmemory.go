package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	for i, byt := range arr {
		hexEq := convertHex(int(byt))
		withoutanewline(hexEq)

		if (i+1)%4 == 0 && i != 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

	}
	for _, ch := range arr {
		if ch >= 33 && ch <= 127 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func convertHex(nb int) string {
	if nb == 0 {
		return "00"
	}
	hexdigits := "0123456789abcdef"
	runes := []rune{}

	for nb > 0 {
		runes = append(runes, rune(hexdigits[nb%16]))
		nb /= 16
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func withoutanewline(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	num := 10

	hex := convertHex(num)
	fmt.Println(hex)
}
