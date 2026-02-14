package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	runes := []rune(os.Args[1])
	from := []rune(os.Args[2])[0] // can also add length check for empty args
	to := []rune(os.Args[3])[0]
	for i, v := range runes {
		if v == from {
			runes[i] = to
		}
	}
	fmt.Println(string(runes))
}
