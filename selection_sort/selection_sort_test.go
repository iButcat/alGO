package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	var myArray = []int{3, 2, 8, 5, 15, 12, 100, 56}
	SelectionSort(myArray)
	assert.Equal(t, myArray, []int{2, 3, 5, 8, 12, 15, 56, 100})
}
