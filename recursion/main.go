package main

import "fmt"

func main() {
	var mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sum(mySlice))

	fmt.Println(counter(mySlice))

	fmt.Println(maximum(mySlice))
}
