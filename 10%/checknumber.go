package main

func CheckNumber(arg string) bool {
	for _, v := range arg {
		if isNum(v) {
			return true
		}
	}
	return false
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}
