package main

func CountAlpha(s string) int {
	count := 0
	for _, v := range s {
		if isAlpha(v) {
			count++
		}
	}
	return count
}

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}
