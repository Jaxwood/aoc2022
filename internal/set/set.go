package set

type Set struct {
	Items map[rune]bool
}

func ToSet(str string) Set {
	result := map[rune]bool{}
	for _, s := range str {
		result[s] = true
	}
	return Set{result}
}

func (s Set) Get(key rune) bool {
  _, ok := s.Items[key]
  return ok
}

func (s Set) Len() int {
	return len(s.Items)
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

func (s Set) Difference(other Set) Set {
	diff := map[rune]bool{}
	for k, _ := range s.Items {
		_, ok := other.Items[k]
		if !ok {
			diff[k] = true
		}
	}
	return Set{diff}
}
