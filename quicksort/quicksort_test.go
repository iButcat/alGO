package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestQuickSortBaseCase(t *testing.T) {
	var mySlice = []int{2}
	res := quicksort(mySlice, 0, len(mySlice)-1)
	assert.Equal(t, res, []int{2})
}

func TestQuickSort(t *testing.T) {
	var mySlice = []int{6, -10, 7, 12, 10, 200, 0, -3, 15, -2220, 344}
	res := quicksort(mySlice, 0, len(mySlice)-1)

	assert.Equal(t, res, []int{-2220, -10, -3, 0, 6, 7, 10, 12, 15, 200, 344})
}
