package main

import "github.com/01-edu/z01"

func PrintRevComb() {
	for i := '9'; i >= '2'; i-- {
		for k := i - 1; k >= '1'; k-- {
			for m := k - 1; m >= '0'; m-- {
				z01.PrintRune(i)
				z01.PrintRune(k)
				z01.PrintRune(m)

				if !(i == '2' && k == '1' && m == '0') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
