package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	for _, r := range arg {
		if r >= '1' && r <= '9' {
			return true
		}
	}
	return false
}
