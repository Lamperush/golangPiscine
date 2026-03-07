package main

import (
	"fmt"
	"os"
)

func AddPrimeSum() {
	if len(os.Args) != 2 || BaseAtoi(os.Args[1]) < 2 {
		fmt.Println(0)
		return
	}
	nb := BaseAtoi(os.Args[1])
	result := 0
	for i := nb; i > 1; i-- {
		if isPrime(i) {
			result += i
		}
	}
	fmt.Println(result)
}
func isPrime(nb int) bool {
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func BaseAtoi(str string) int { // this function expects strings that contain only numbers
	runes := []rune(str)
	result := 0
	for _, v := range runes {
		result = result*10 + int(v-'0')
	}
	return result
}
