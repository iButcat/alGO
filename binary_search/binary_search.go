package main

import "fmt"

func BinarySearch(sortedList []int, element int) int {
	var low int = 0
	var high int = len(sortedList) - 1

	for low <= high {
		var middle int = low + (high-low)/2
		var guess int = sortedList[middle]
		if guess == element {
			return middle
		} else if guess > element { //guess too high
			high = middle - 1
		} else { // guess too low
			low = middle + 1
		}
	}
	return 0
}

func main() {
	var mySlice []int
	mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var firstElement int = 10
	position := BinarySearch(mySlice, firstElement)
	fmt.Println("Nb:", firstElement, "is at the position:", position)
}
