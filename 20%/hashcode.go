package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

//	func main() {
//		fmt.Println(HashCode("A"))
//		fmt.Println(HashCode("AB"))
//		fmt.Println(HashCode("BAC"))
//		fmt.Println(HashCode("Hello World"))
//	}
func main() {
	table := []string{"Z", "Hi!", "BB198365", "sabito", "14 Avril 1912", "zyx987bca", "		pool-2020", "965truma747", " Mercedes-AMG GT"}
	for _, s := range table {
		challenge.Function("HashCode", HashCode, solutions.HashCode, s)
	}
}
func HashCode(dec string) string {
	runes := []rune(dec)
	size := rune(len(dec))
	for i := range runes {
		runes[i] = (runes[i] + size) % 127
		if runes[i] < 33 {
			runes[i] += 33
		}
	}
	return string(runes)
}
