package main

import (
    "fmt"
)

func FromTo(from int, to int) string {
    if from < 0 || from > 99 || to < 0 || to > 99 {
        return "Invalid\n"
    }

    result := ""
    if from <= to {
        for i := from; i <= to; i++ {
            if i < 10 {
                result += "0" + string('0'+i)
            } else {
                tens := i / 10
                ones := i % 10
                result += string('0'+tens) + string('0'+ones)
            }

            if i != to {
                result += ", "
            }
        }
    } else {
        for i := from; i >= to; i-- {
            if i < 10 {
                result += "0" + string('0'+i)
            } else {
                tens := i / 10
                ones := i % 10
                result += string('0'+tens) + string('0'+ones)
            }

            if i != to {
                result += ", "
            }
        }
    }
    result += "\n"

    return result
}

func main() {
    fmt.Print(FromTo(1, 10))
    fmt.Print(FromTo(10, 1))
    fmt.Print(FromTo(10, 10))
    fmt.Print(FromTo(100, 10))
}