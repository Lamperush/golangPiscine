package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	var chunks [][]int
	for i := 0; i < len(slice); i = i + size { // interval
		end := i + size       //separator based on the size
		if end > len(slice) { // if "end" is out of bounds, then copy up to max len
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	fmt.Println(chunks)
}
