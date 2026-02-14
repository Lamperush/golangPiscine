package main

func ZipString(s string) string {
	runes := []rune(s)
	result := ""
	counter := 1
	for i := range runes {
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			counter++
		} else {
			result += Itoa(counter) + string(runes[i])
			counter = 1
		}
	}
	return result
}
