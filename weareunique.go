package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))       // Expect 2: 'f' and 'b' are unique
	fmt.Println(WeAreUnique("", ""))             // Expect -1: both strings empty
	fmt.Println(WeAreUnique("abc", "def"))       // Expect 6: all characters are unique
	fmt.Println(WeAreUnique("abc", "abc"))       // Expect 0: all common
	fmt.Println(WeAreUnique("abcd", "cdef"))     // Expect 4: 'a','b','e','f' are unique
	fmt.Println(WeAreUnique("aabbcc", "ddeeff")) // Expect 6: no common characters
	fmt.Println(WeAreUnique("aaa", "aaa"))       // Expect 0: all common, one unique char
	fmt.Println(WeAreUnique("a", ""))            // Expect 1
	fmt.Println(WeAreUnique("", "z"))            // Expect 1
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	// Sets of unique characters
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	for _, ch := range str1 {
		set1[ch] = true
	}
	for _, ch := range str2 {
		set2[ch] = true
	}

	uniqueCount := 0

	// Count characters unique to set1
	for ch := range set1 {
		if !set2[ch] {
			uniqueCount++
		}
	}

	// Count characters unique to set2
	for ch := range set2 {
		if !set1[ch] {
			uniqueCount++
		}
	}

	return uniqueCount
}
