package main

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}
	result := ""
	runes := []rune(str)

	for i := 2; i < len(runes); i = i + 3 {
		result += string(runes[i])
	}
	return result + "\n"
}
