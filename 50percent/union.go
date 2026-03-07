package main

import (
	"fmt"
	"os"
)

func Union() {
	if len(os.Args) != 3 {
		return
	}
	strs := os.Args[1] + os.Args[2]
	var result []rune
	inResult := make(map[byte]bool)

	for i := 0; i < len(strs); i++ {
		if !inResult[strs[i]] { // if the character does not exist inside the result then append
			result = append(result, rune(strs[i]))
			inResult[strs[i]] = true // make a map of resulting string
		}
	}
	fmt.Println(string(result))
}
