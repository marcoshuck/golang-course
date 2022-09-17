package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Non-empty array A consisting of N ints is given:
// First node (head) located at index 0;
// the value of a node located at index K is A[K]:
// if the value of a node is -1 then it is the last node of the list
// otherwise, the successor of a node located at index K is located at index A[K] (you can assume that A[K] is a valid index, that is 0 <= A[K] < N)

// 1. Read the whole problem
// 2. Identify the inputs
// 	  [1, 4, -1, 3, 2]
// 3. What's the output?
//	  length of resultant array: 4.
// 5. Can I create more test cases? How?
//	  Add inputs with different values and lengths.
func TestLengthOutput(t *testing.T) {
	var input []int
	var output int
	assert.IsType(t, output, Length(input))
}

func TestLengthValue(t *testing.T) {
	a := []int{1, 4, -1, 3, 2}
	assert.Equal(t, 4, Length(a))

	b := []int{1, 3, -1, 2, 4}
	assert.Equal(t, 4, Length(b))

	c := []int{1, 3, -1, 2, 0, 0, 0}
	assert.Equal(t, 4, Length(c))

	d := []int{2, 1, -1}
	assert.Equal(t, 2, Length(d))
}

// do-while loop exercise
func Length(input []int) int {
	var result []int

	// Getting the first value, that will have the next index
	next := input[0]

	// If next is already -1, we should exit. If not, we should run a certain set of instructions.
	for next != -1 {
		// Save the number you found
		result = append(result, next)

		// Get the next element
		next = input[next]
	}

	result = append(result, -1)

	return len(result)
}
