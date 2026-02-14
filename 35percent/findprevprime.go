package main

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if nb == 2 {
		return nb
	}
	for i := nb; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
