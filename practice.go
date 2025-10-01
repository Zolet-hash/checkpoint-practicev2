package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	if n <= 0 {
		fmt.Println("0")
		return
	}

	sum := 0
	for i := 2; i <= n; i++ {
		p := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				p = false
				break
			}
		}
		if p {
			sum += i
		}
	}
	fmt.Println(sum)
}
