package main

import (
	"fmt"
)

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	length := len([]rune(str))
	runes := []rune(str)
	var thirdChar []rune

	for i := 2; i < length; i += 3 {
		thirdChar = append(thirdChar, runes[i])
	}
	return string(thirdChar) + "\n"
}
