package main

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
