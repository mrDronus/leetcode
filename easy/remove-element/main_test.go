package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	a := []int{3, 2, 2, 3}
	k := removeElement(a, 3)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{2, 2}, a[:k])

	a = []int{0, 1, 2, 2, 3, 0, 4, 2}
	k = removeElement(a, 2)
	assert.Equal(t, 5, k)
	s := a[:k]
	slices.Sort(s)
	assert.Equal(t, []int{0, 0, 1, 3, 4}, s)
}
