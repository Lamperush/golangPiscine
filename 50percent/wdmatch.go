package main

import (
	"fmt"
	"os"
)

func main() {
	Wdmatch()
}

func Wdmatch() {
	if len(os.Args) < 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	k := 0
	for i := 0; i < len(s2) && k < len(s1); i++ { // loop goes over s2 and if s1 found in s2 then it stops the loop
		if s1[k] == s2[i] { // if character in s1 found then move to the next character
			k++
		}
	}
	if k == len(s1) {
		fmt.Println(s1)
		return
	}
}
