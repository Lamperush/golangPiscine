package main

import (
	"strings"
)

func SaveAndMiss(arg string, num int) string {
	if num < 1 || arg == "" {
		return arg
	}
	var strs []string
	flag := true
	for i := 0; i < len(arg); i = i + num {
		if flag {
			end := i + num      //separator based on the size
			if end > len(arg) { // if "end" is out of bounds, then copy up to max len
				end = len(arg)
			}
			strs = append(strs, arg[i:end])
		}
		flag = !flag // alternate the flag, so it jumps after appending
	}
	result := strings.Join(strs, "")
	return result
}
