package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	n := []int{1, 2, 3, 0, 0, 0}
	merge(n, 3, []int{2, 5, 6}, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, n)

	n = []int{4, 5, 6, 0, 0, 0}
	merge(n, 3, []int{1, 2, 3}, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, n)
}
