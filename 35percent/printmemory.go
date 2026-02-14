package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	for i := range arr {
		first := arr[i] / 16
		second := arr[i] % 16
		z01.PrintRune(rune(base[first]))
		z01.PrintRune(rune(base[second]))
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for i := range arr {
		if arr[i] < 33 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(arr[i]))
		}
	}
	z01.PrintRune('\n')
}
