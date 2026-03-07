package main

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}

	size := len(slice1) + len(slice2)
	result := make([]int, size)
	var short, long []int

	if len(slice1) >= len(slice2) {
		short, long = slice2, slice1
	} else {
		short, long = slice1, slice2
	}

	index := 0
	for i := range short { // copying up to the length of the shorter slice
		result[index] = long[i]
		index++
		result[index] = short[i]
		index++
	}
	copy(result[index:], long[len(short):]) // copying the remaining of the longer slice
	return result
}
