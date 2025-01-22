package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	v := []int{1, 1, 2}
	k := removeDuplicates(v)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{1, 2}, v[:k])

	v = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k = removeDuplicates(v)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, v[:k])
}
