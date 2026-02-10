package main

// euclidean algorithm
func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b // 10 % 40 = the remainder is actually 10,
		// so the numbers will swap
	}
	return a
}

// my original code below

// func Gcd(a, b int) int {
// 	if a < 0 || b < 0 {
// 		return 0
// 	}
// 	for a != 0 && b != 0 {
// 		if a > b {
// 			a = a % b
// 		} else {
// 			b = b % a
// 		}
// 	}
// 	if a < b {
// 		a = b
// 	}
// 	return a
// }
