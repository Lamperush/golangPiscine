package main

func LastWord(s string) string {
	if s == "" {
		return "\n"
	}
	runes := []rune(s)
	start := 0
	end := 0
	isLast := true
	for i := len(runes) - 1; i >= 0; i-- {
		if isLast && !(runes[i] == ' ' || runes[i] == '\t' || runes[i] == '\n') {
			end = i + 1 // because [:end] it cuts up to, w/t inclusion
			isLast = false
		}
		if !isLast && (runes[i] == ' ' || runes[i] == '\t' || runes[i] == '\n') {
			start = i + 1
			break
		}
	}
	return string(runes[start:end]) + "\n" // I could've done s[] but
	//  it will be byte manipulation and can lead to undefined behaviour
}
