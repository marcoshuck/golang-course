package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var result []int

	// Test case: 9
	result = twoSum([]int{2, 7, 11, 15}, 9)

	assert.Len(t, result, 2)
	assert.Contains(t, result, 0)
	assert.Contains(t, result, 1)

	// Test case: 18
	result = twoSum([]int{2, 7, 11, 15}, 18)

	assert.Len(t, result, 2)
	assert.Contains(t, result, 1)
	assert.Contains(t, result, 2)

	// Test case: 13
	result = twoSum([]int{2, 7, 11, 15}, 13)

	assert.Len(t, result, 2)
	assert.Contains(t, result, 0)
	assert.Contains(t, result, 2)

	// Test case: 26
	result = twoSum([]int{2, 7, 11, 15}, 26)

	assert.Len(t, result, 2)
	assert.Contains(t, result, 2)
	assert.Contains(t, result, 3)
}

// [2, 7, 11, 15] O(2*n
// [2, 7, 11, 15]
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		first := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if first+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
