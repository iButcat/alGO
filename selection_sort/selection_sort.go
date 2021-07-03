package main

import "fmt"

func SelectionSort(array []int) {
	var lenght int = len(array)
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ { //next position
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	var myArray = []int{10, 5, 3, 7, 18, 2, 4}
	SelectionSort(myArray)
	fmt.Println(myArray)
}
