package main

import (
    "fmt"
)

func Itoa(n int) string {
    if n == 0 {
        return "0"
    }

    isNegative := false
    if n < 0 {
        isNegative = true
        n = -n
    }

    var result []rune
    for n > 0 {
        digit := n % 10
        result = append(result, rune(digit+'0'))
        n /= 10
    }

    if isNegative {
        result = append(result, '-')
    }

    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}