package main

func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	if from == to {
		if from < 10 {
			return "0" + Itoa2(from) + "\n"
		}
		return Itoa2(from) + "\n"
	}
	result := ""
	if to > from {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + Itoa2(i)
			} else {
				result += Itoa2(i)
			}
			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + Itoa2(i)
			} else {
				result += Itoa2(i)
			}
			if i != to {
				result += ", "
			}
		}
	}
	return result + "\n"
}
func Itoa2(nb int) string {
	if nb < 10 {
		return string(rune(nb + '0'))
	}
	return Itoa2(nb/10) + string(rune(nb%10+'0'))
}

// Gemini way
// func FromTo2(from int, to int) string {
// 	if from < 0 || from > 99 || to < 0 || to > 99 {
// 		return "Invalid\n"
// 	}
// 	// 2. Determine Step (1 or -1)
// 	step := 1
// 	if from > to {
// 		step = -1
// 	}

// 	result := ""
// 	for {
// 		numStr := Itoa(from)
// 		if from < 10 {
// 			numStr = "0" + numStr
// 		}
// 		result += numStr
// 		if from == to {
// 			break
// 		}
// 		result += ", "
// 		from += step
// 	}

// 	return result + "\n"
// }
