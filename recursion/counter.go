package main

func counter(mySlice []int) int {
	if len(mySlice) == 0 {
		return 0
	}
	return 1 + counter(mySlice[1:])
}
