package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))
}

func CanJump(slice []uint) bool {
	if len(slice) == 0 {
		return false
	}
	if len(slice) == 1 {
		return true
	}

	pos := 0
	for {
		step := int(slice[pos]) //index
		pos += step

		if pos == len(slice)-1 {
			return true
		}
		if pos >= len(slice) {
			return false
		}
	}

}


Approach

You can simulate the jumps step by step:

Start at index 0.

At each position, jump forward exactly the number of steps indicated by the value at that position.

If you reach the last index exactly → return true.

If you jump outside the slice or get stuck → return false.

Explanation:

If the slice is empty, return false.

If the slice has 1 element, return true (already at the last index).

Set pos to 0.

Keep jumping:

pos += steps where steps = value at current position.

If pos equals last index → success.

If pos goes beyond the last index → fail.


