package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
	assert.Equal(t, "0", addBinary("0", "0"))
	assert.Equal(t, "1000001", addBinary("1000000", "1"))
	assert.Equal(t, "1000", addBinary("1", "111"))
}
