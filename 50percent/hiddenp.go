package main

import (
	"fmt"
	"os"
)

func HiddenP() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if len(s1) == 0 {
		fmt.Println("1")
		return
	}
	index := 0
	for k := 0; index < len(s1) && k < len(s2); k++ {
		if s1[index] == s2[k] {
			index++ // if we find a match, then move the index to look for the next symbol
		}
	}
	if index == len(s1) {
		fmt.Println("1")
		return
	}
	fmt.Println("0")
}
