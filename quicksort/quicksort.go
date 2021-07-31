package main

import (
	"fmt"
)

func partition(mySlice []int, start, end int) int {
	var pivot = mySlice[start]
	var leftwall = start

	for i := start + 1; i < end; i++ {
		if mySlice[i] < pivot {
			leftwall++
			mySlice[i], mySlice[leftwall] = mySlice[leftwall], mySlice[i]
		}
	}

	mySlice[start], mySlice[leftwall] = mySlice[leftwall], mySlice[start]

	return leftwall
}

func quicksort(mySlice []int, start, end int) []int {
	// Base case
	if len(mySlice) == 1 {
		return mySlice
	}

	if start < end {
		var pivot = partition(mySlice, start, end)
		quicksort(mySlice, start, pivot)
		quicksort(mySlice, pivot+1, end)
		return mySlice
	}
	return mySlice
}

func main() {
	var mySlice = []int{6, -10, 7, 12, -100, 10, 200, 0, -3, 15, -2220, 344, 500}
	fmt.Println(quicksort(mySlice, 0, len(mySlice)-1))
}
