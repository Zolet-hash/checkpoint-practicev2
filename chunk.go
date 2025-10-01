package main

import (
	"fmt"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println() //print empty slice
		return
	}

	var result [][]int //two-dimensional slice to hold chunks
	for i := 0; i < len(slice); i += size { //iterate in steps of 'size'
		end := i + size //calculate end index
		if end > len(slice) { //a
			end = len(slice) //adjust end if it exceeds slice length
		}
		result = append(result, slice[i:end]) 
	}
	fmt.Println(result)
}
