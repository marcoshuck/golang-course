package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(123))
}

func TestSplit(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, split(123))
	assert.Equal(t, []int{1, 2, 2, 1}, split(1221))
	assert.Equal(t, []int{1, 2, 1}, split(121))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, reverse([]int{1, 2, 3}))
}

func reverse(arr []int) []int {
	result := make([]int, len(arr))
	var count int
	for i := len(arr) - 1; i >= 0; i-- {
		result[count] = arr[i]
		count++
	}
	return result
}

func split(x int) []int {
	var result []int
	for x >= 10 {
		result = append(result, x%10)
		x = x / 10
	}
	return reverse(append(result, x))
}

func isPalindrome(x int) bool {
	digits := split(x)
	reversed := reverse(digits)
	for i := 0; i < len(digits); i++ {
		if digits[i] != reversed[i] {
			return false
		}
	}
	return true
}

var romans = map[int]string{
	10: "X",
	9:  "IX",
}
