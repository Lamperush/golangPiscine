package main

import (
	"fmt"
	"os"
)

func Reversestrcap() {
	if len(os.Args) < 2 {
		return
	}

	strs := os.Args[1:]

	for _, str := range strs {
		result := make([]rune, 0, len(str)) // initialize empty slice of runes with length of str
		last := len(str) - 1
		for i, v := range str {
			if i < last && rune(str[i+1]) == ' ' || i == last {
				result = append(result, toUpper(v))
			} else {
				result = append(result, toLower(v))
			}
		}
		fmt.Println(string(result))
	}
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}
