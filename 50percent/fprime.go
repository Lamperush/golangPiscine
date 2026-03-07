package main

import (
	"fmt"
	"os"
)

func Fprime() {
	if len(os.Args) != 2 {
		return
	}
	nb := Atoi(os.Args[1])
	if nb == 0 {
		return
	}
	if nb == 1 {
		fmt.Println()
		return
	}
	for i := 2; i <= nb; i++ {
		for nb%i == 0 {
			fmt.Printf("%d", i)
			nb = nb / i
			if nb > 1 {
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
}
func Atoi(str string) int {
	runes := []rune(str)
	result := 0
	for _, v := range runes {
		if v < '0' || v > '9' {
			return 0
		}
		result = result*10 + int(v-'0')
	}
	return result
}
