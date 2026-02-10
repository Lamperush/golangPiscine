package main

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	result := ""
	if isUpper(runes[len(runes)-1]) { // checking the last letter for uppercase
		return s
	}
	for i, v := range runes {
		if !isLetter(v) || isUpper(v) && isUpper(runes[i+1]) { // letter check
			return s // checking for consecutive uppercase
		}
		if isUpper(v) && i > 0 { // if not first letter and uppercase
			result += "_" + string(v)
		} else {
			result += string(v)
		}
	}
	return result
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
func isLetter(r rune) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}
