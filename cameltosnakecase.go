package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	for i, r := range s {
		if r >= '0' && r <= '9' {
			return s
		}
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return s
		}
		if i > 0 && r >= 'A' && r <= 'Z' {
			prev := rune(i-1)

			if prev >= 'A' && prev <= 'Z' {
				return s
			}
		}
		last := rune(s[len(s)-1])

		if last >= 'A' && last <= 'Z' {
			return s
		}
	}

	result := []rune{}
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(result)
}
