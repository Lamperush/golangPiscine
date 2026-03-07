package main

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}
	size := len(slice1) + len(slice2)
	result := make([]int, size)
	copy(result, slice1)
	copy(result[len(slice1):], slice2)
	return result
}
