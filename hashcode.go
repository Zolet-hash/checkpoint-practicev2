package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	length := len(dec)
	var hashedstr string

	for _, r := range dec {
		hashedval := (int(r) + length) %127

		if hashedval < 32 {
			hashedval += 33
		}

		hashedstr += string(rune(hashedval))
	}
	return hashedstr
}
