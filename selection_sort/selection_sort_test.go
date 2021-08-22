package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	var myArray = []int{3, 2, 8, 5, 15, 12, 100, 56}
	SelectionSort(myArray)
	assert.Equal(t, myArray, []int{2, 3, 5, 8, 12, 15, 56, 100})
}

func TestSelectionSortBiggerSlice(t *testing.T) {
	var firstSlice = []int{3, 2, 8, 5, 15, 12, 100, 56, 12, 1, 7, 1000,
		123487, 3400, 223, 123, 80032, 23, 129, -10}
	var secondSlice = []int{3, 2, 8, 5, 15, 12, 100, 56, 12, 1, 7, 1000,
		123487, 3400, 223, 123, 80032, 23, 129, -10}
	SelectionSort(firstSlice)
	sort.Ints(secondSlice)
	assert.Equal(t, firstSlice, secondSlice)
}

func BenchmarkSelectionSortSmallerSlice(b *testing.B) {
	var myArray = []int{3, 2, 8, 5, 15, 12, 100, 56}
	for i := 0; i < b.N; i++ {
		SelectionSort(myArray)
	}
}

func BenchmarkSelectionSortBiggerSlice(b *testing.B) {
	var myArray = []int{3, 2, 8, 5, 15, 12, 100, 56, 12, 1, 7, 1000,
		123487, 3400, 223, 123, 80032, 23, 129, -10}
	for i := 0; i < b.N; i++ {
		SelectionSort(myArray)
	}
}
