package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5)) // 5
	fmt.Println(FindPrevPrime(4)) // 3
	fmt.Println(FindPrevPrime(1)) // 0
}

func FindPrevPrime(nb int) int {
	for i := nb; i > 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
