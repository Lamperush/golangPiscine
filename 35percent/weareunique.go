package main

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	for _, v := range str1 {
		m1[v] = true
	}

	for _, v := range str2 {
		m2[v] = true
	}
	counter := 0
	for i := range m1 {
		if !m2[i] {
			counter++
		}
	}
	for i := range m2 {
		if !m1[i] {
			counter++
		}
	}
	return counter
}
