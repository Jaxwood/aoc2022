package maps

type Set struct {
	Items map[rune]bool
}

func StrToMap(str string) map[rune]bool {
	result := map[rune]bool{}
	for _, s := range str {
		result[s] = true
	}
	return result
}

func (s Set) Intersect(other Set) Set {
	same := map[rune]bool{}

	for k, _ := range s.Items {
		for kv, _ := range other.Items {
			if k == kv {
				same[k] = true
			}
		}
	}

	return Set{same}
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
