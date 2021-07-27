package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterBaseCase(t *testing.T) {
	var mySlice []int
	res := counter(mySlice)

	assert.Equal(t, res, 0)
}

func TestCounter(t *testing.T) {
	var mySlice = []int{1, 2, 3, 4, 5, 6, 7}
	res := counter(mySlice)

	assert.Equal(t, res, 7)
}
