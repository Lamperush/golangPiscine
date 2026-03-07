package main

import (
	"fmt"
	"os"
)

func Inter() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	inS2 := make(map[byte]bool)
	for i := 0; i < len(s2); i++ {
		inS2[s2[i]] = true // we get a map of all the s2 characters
	}

	var result []rune
	inResult := make(map[byte]bool)
	for i := 0; i < len(s1); i++ {
		if inS2[s1[i]] && !inResult[s1[i]] { // if s1 charachter exists in s2 and does not in the result
			result = append(result, rune(s1[i]))
			inResult[s1[i]] = true // make a map of resulting string
		}
	}
	fmt.Println(string(result))
}
