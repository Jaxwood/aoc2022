package main

func day06a(file string) int {
	for i := 3; i < len(file); i++ {
		m := map[string]bool{}
		m[string(file[i])] = true
		m[string(file[i-1])] = true
		m[string(file[i-2])] = true
		m[string(file[i-3])] = true

		if len(m) == 4 {
			return i + 1
		}
	}
	return 0
}

func day06b(file string) int {
	for i := 13; i < len(file); i++ {
		m := map[string]bool{}
		m[string(file[i])] = true
		m[string(file[i-1])] = true
		m[string(file[i-2])] = true
		m[string(file[i-3])] = true
		m[string(file[i-4])] = true
		m[string(file[i-5])] = true
		m[string(file[i-6])] = true
		m[string(file[i-7])] = true
		m[string(file[i-8])] = true
		m[string(file[i-9])] = true
		m[string(file[i-10])] = true
		m[string(file[i-11])] = true
		m[string(file[i-12])] = true
		m[string(file[i-13])] = true

		if len(m) == 14 {
			return i + 1
		}
	}
	return 0
}
