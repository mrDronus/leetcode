package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
	assert.Equal(t, 6, strStr("ssdbutsad", "sad"))
	assert.Equal(t, 4, strStr("mississippi", "issip"))
}
