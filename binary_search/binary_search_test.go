package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchWithElementsExist(t *testing.T) {
	var mySortedList = []int{10, 20, 30, 40, 50, 60}
	var firstLookingFor int = 20
	var res = BinarySearch(mySortedList, firstLookingFor)

	var secondLookingFor int = 60
	var resOne = BinarySearch(mySortedList, secondLookingFor)

	assert.Equal(t, res, 1)
	assert.Equal(t, resOne, 5)
}

func TestBinarySearchWithElementNotFound(t *testing.T) {
	var mySortedList = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var elementDoesNotExist = 10
	var res = BinarySearch(mySortedList, elementDoesNotExist)

	assert.Equal(t, res, -1)
}
