package main

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	if isLower(s[0]) {
		return false
	}
	for i, v := range s {
		if i < len(s)-1 && v == ' ' && isLower(s[i+1]) {
			return false
		}
	}
	return true
}

func isLower(r byte) bool {
	return r >= 'a' && r <= 'z'
}
