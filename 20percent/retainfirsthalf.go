package main

func RetainFirstHalf(str string) string {
	if str == "" || len(str) == 1 {
		return str
	}
	return str[:(len(str) / 2)]
}
