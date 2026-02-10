package main

func CountChar(str string, c rune) int {
	if str == "" {
		return 0
	}
	match := 0

	for _, v := range str {
		if v == c {
			match++
		}
	}
	return match
}
