package maps

func toMap(str string) map[rune]bool {
	result := map[rune]bool{}
	for _, s := range str {
		result[s] = true
	}
	return result
}

