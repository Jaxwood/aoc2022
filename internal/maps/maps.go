package maps

func StrToMap(str string) map[rune]bool {
	result := map[rune]bool{}
	for _, s := range str {
		result[s] = true
	}
	return result
}

func Union(maps []map[rune]bool) []rune {
	first := maps[0]
	rest := maps[1:]
	for _, m := range rest {
		first = union([]map[rune]bool{first, m})
	}
	return keys(first)
}

func Difference(first map[int]bool, second map[int]bool) map[int]bool {
	diff := map[int]bool{}
	for k, _ := range first {
		_, ok := second[k]
		if !ok {
			diff[k] = true
		}
	}
	return diff
}

func keys(candidate map[rune]bool) []rune {
	var result []rune
	for k, _ := range candidate {
		result = append(result, k)
	}
	return result
}

func union(maps []map[rune]bool) map[rune]bool {
	first := maps[0]
	second := maps[1]
	same := []rune{}

	for k, _ := range first {
		for kv, _ := range second {
			if k == kv {
				same = append(same, k)
			}
		}
	}

	return StrToMap(string(same))
}
