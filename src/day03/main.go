package main

import (
	"strings"
)

func day03a(filename string) int {
	lines := strings.Split(filename, "\n")
	sum := 0
	scores := map[rune]int{}
	// lowercases
	for i := 97; i < 97+26; i++ {
		scores[rune(i)] = i - 96
	}
	// uppercases
	for i := 65; i < 65+26; i++ {
		scores[rune(i)] = 27 + (i - 65)
	}
	for _, line := range lines {
		chars := map[rune]bool{}
		start := line[:len(line) / 2]
		end := line[len(line) / 2:]
		for _, c := range start {
			chars[c] = true
		}
		for _, c := range end {
			_, ok := chars[c]
			if ok {
				sum += scores[c]
				break
			}
		}
	}
	return sum
}
