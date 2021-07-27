package main

func maximum(mySlice []int) int {
	if len(mySlice) == 2 {
		if mySlice[0] > mySlice[1] {
			return mySlice[0]
		} else {
			return mySlice[1]
		}
	}
	var max_nb = maximum(mySlice[1:])
	if mySlice[0] > max_nb {
		return mySlice[0]
	} else {
		return max_nb
	}
}
