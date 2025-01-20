package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
	assert.Equal(t, []int{0, 2}, twoSum([]int{-3, 4, 3, 90}, 0))
}

func TestTwoSum2(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum2([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum2([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum2([]int{3, 3}, 6))
	assert.Equal(t, []int{0, 2}, twoSum2([]int{-3, 4, 3, 90}, 0))
}
