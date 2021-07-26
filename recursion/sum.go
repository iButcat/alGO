package main

import (
	"fmt"
)

func sum(mySlice []int) int {
	if len(mySlice) == 0 {
		return 0
	}
	return mySlice[0] + sum(mySlice[1:])
}

func main() {
	var mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(mySlice))
}
