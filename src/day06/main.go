package main

func day06(file string, size int) int {
	for i := size - 1; i < len(file); i++ {
		m := map[string]bool{}
		start := -1*size + 1
		for j := 0; j >= start; j-- {
			m[string(file[j+i])] = true
		}

		if len(m) == size {
			return i + 1
		}
	}
	return 0
}
