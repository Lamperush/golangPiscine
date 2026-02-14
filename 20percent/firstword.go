package main

func FirstWord(s string) string {
	runes := []rune(s)
	result := ""
	for _, v := range runes {
		if isPrintable(v) {
			result += string(v)
		}
		if (v == ' ' || v == '\n' || v == '\t') && result != "" {
			break
		}
	}
	return result + "\n"
}

func isPrintable(r rune) bool {
	return r > 32
}
