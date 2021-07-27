package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBaseCase(t *testing.T) {
	var mySlice = []int{1, 2}
	res := maximum(mySlice)

	assert.Equal(t, res, 2)
}

func TestMaximumRecursion(t *testing.T) {
	var mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := maximum(mySlice)

	assert.Equal(t, res, 10)
}
