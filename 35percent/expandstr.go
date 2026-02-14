package main

import (
	"fmt"
	"os"
)

func ExpandStr() { // should rename for it to work
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println()
		return
	}
	runes := []rune(os.Args[1])
	result := ""
	needsSpace := false
	for _, v := range runes {
		if !(v == ' ' || v == '\t' || v == '\n') {
			if needsSpace {
				result += "   "
				needsSpace = false
			}
			result += string(v)
		} else {
			if result != "" {
				needsSpace = true
			}
		}
	}
	fmt.Println(result)
}
