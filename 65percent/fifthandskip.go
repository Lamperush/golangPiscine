package main

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	var result []rune
	skip := false
	count := 0
	for i, v := range str {
		if v == ' ' {
			continue
		} else if skip { // skip the 6th non-space character
			skip = false
			continue
		}
		result = append(result, v)
		count++
		if count%5 == 0 && i < len(str)-1 { // putting space after fifth char but not on the last one
			result = append(result, ' ')
			skip = true
		}
	}
	return string(result) + "\n"
}
