package main

func sum(mySlice []int) int {
	if len(mySlice) == 0 {
		return 0
	}
	return mySlice[0] + sum(mySlice[1:])
}
