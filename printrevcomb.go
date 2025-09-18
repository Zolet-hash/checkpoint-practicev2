package main

import (
	"github.com/01-edu/z01"
)

func main() {
	firstComb := true

	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				if i > j && j > k {
					if !firstComb {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					firstComb = false
				}
			}
		}
	}
	z01.PrintRune('\n')
}
