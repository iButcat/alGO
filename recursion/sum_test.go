package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := sum(mySlice)

	assert.Equal(t, res, 45)
}
