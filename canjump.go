package main

import (
	"fmt"

)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // false

	input3 := []uint{0}
	fmt.Println(CanJump(input3)) // true
}

func CanJump(arr []uint) bool {
    if len(arr) == 0 {
        return false
    }
    if len(arr) == 1 {
        return true
    }
    visited := make(map[int]bool) //to track visited positions to avoid infinite loops
    pos := 0 //starting position
    
    for { 
        if visited[pos] { 
            return false 
        }
        
        visited[pos] = true 	//mark current position as visited
    
        jump := int(arr[pos]) //number of steps to jump
        next := pos + jump //calculate next position
        
        if next == len(arr) -1 {
            return true //landed exactly in the last index
        }
        if next >= len(arr) || jump == 0 {
            return false //jumped out of bounds or stuck
        }
        pos = next //move to next position
    }
    
}
// CanJump determines if you can reach the last index of the slice
// by jumping according to the values in the slice.
// Each value represents the exact number of steps to jump forward from that position.
// You start at index 0.

// Examples:

// Input: []uint{2, 3, 1, 1, 4}
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Input: []uint{3, 2, 1, 0, 4}
// Output: false
// Explanation: You will always arrive at index 3 where the value is 0. 
// You cannot move forward from there.

// Input: []uint{0}
// Output: true
// Explanation: You are already at the last index.

// Constraints:

// The input slice can be empty or contain up to 10^5 elements.
// Each element in the slice is a non-negative integer (0 ≤ arr[i] ≤ 10^5).		

// Approach

// You can simulate the jumps step by step:

// Start at index 0.

// At each position, jump forward exactly the number of steps indicated by the value at that position.

// If you reach the last index exactly → return true.

// If you jump outside the slice or get stuck → return false.
// Explanation:

// If the slice is empty, return false.

// If the slice has 1 element, return true (already at the last index).

// Set pos to 0.

// Keep jumping:

// pos += steps where steps = value at current position.

// If pos equals last index → success.

// If pos goes beyond the last index → fail.


