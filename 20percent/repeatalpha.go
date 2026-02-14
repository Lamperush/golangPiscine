package main

func RepeatAlpha(s string) string {
	runes := []rune(s)
	result := ""

	for _, v := range runes {
		if isLower(v) {
			for i := 0; i < int(v-'a')+1; i++ {
				result += string(v)
			}
		} else if isUpper(v) {
			for i := 0; i < int(v-'A')+1; i++ {
				result += string(v)
			}
		} else {
			result += string(v)
		}
	}
	return result
}

// func isUpper(r rune) bool { // isUpper redeclared in cameltosnakecase
//
//		return r >= 'A' && r <= 'Z'
//	}
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
