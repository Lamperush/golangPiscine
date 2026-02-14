package main

// Signed Itoa
func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	// var result []rune
	var tmp []rune
	sign := true
	if nb > 0 {
		nb = -nb
		sign = false
	}

	for nb != 0 {
		tmp = append(tmp, '0'-rune(nb%10))
		nb /= 10
	}
	// if sign {
	// 	result = append(result, '-')
	// }
	if sign {
		tmp = append(tmp, '-')
	}
	// for i := len(tmp) - 1; i >= 0; i-- {
	// 	result = append(result, tmp[i])
	// }
	for i, k := len(tmp)-1, 0; i > k; i, k = i-1, k+1 {
		tmp[k], tmp[i] = tmp[i], tmp[k]
	}
	return string(tmp)
}

// // "Unsigned" Itoa
// func Itoa(nb int) string {
// 	if nb == 0 {
// 		return "0"
// 	}
// 	if nb == int(-1<<63) { // or I could use this -9223372036854775808
// 		return "-9223372036854775808"
// 	}
// 	var result []rune
// 	if nb < 0 {
// 		result = append(result, '-')
// 		nb *= -1
// 	}
// 	var runes []rune
// 	for nb != 0 {
// 		runes = append(runes, rune(nb%10)+'0')
// 		nb /= 10
// 	}
// 	for i := len(runes) - 1; i >= 0; i-- {
// 		result = append(result, runes[i])
// 	}
// 	return string(result)
// }

// // Recursive Itoa
// func Itoa(nb int) string {
// 	if nb == -1<<63 {
// 		return Itoa(nb/10) + "8"
// 	}
// 	if nb < 0 {
// 		return "-" + Itoa(-nb)
// 	}
// 	if nb < 10 {
// 		return string(rune(nb + '0'))
// 	}
// 	return Itoa(nb/10) + string(rune(nb%10+'0'))
// }
