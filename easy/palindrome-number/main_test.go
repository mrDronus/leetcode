package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(9))
	assert.False(t, isPalindrome(10))
	assert.False(t, isPalindrome(-121))
	assert.True(t, isPalindrome(11))
	assert.True(t, isPalindrome(12321))
	assert.True(t, isPalindrome(1221))
}
