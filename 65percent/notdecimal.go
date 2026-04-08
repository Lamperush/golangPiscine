package main

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	hasDot := false
	for _, v := range dec {
		if v == '.' {
			hasDot = true
		}
		if (v < '0' || v > '9') && v != '.' && v != '-' {
			return dec + "\n" // check for invalid symbol
		}
	}
	if !hasDot {
		return dec + "\n" // if it doesn't have a dot then return itself
	}
	var result []rune
	for _, v := range dec {
		if len(result) == 0 && v == '0' { // check for starting zeroes
			continue
		}
		if v != '.' { //if not a dot, then add it to the result
			result = append(result, v)
		}
	}
	return string(result) + "\n"
}
