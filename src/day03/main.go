package main

import (
	"strings"
	"github.com/jaxwood/aoc2022/maps"
	// "fmt"
	// "golang.org/x/exp/maps"
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
		start := line[:len(line) / 2]
		end := line[len(line) / 2:]
		chars := toMap(start)
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

func day03b(filename string) int {
	// group by lines of 3
	lines := strings.Split(filename, "\n")
	idx := 0
	lineGroup := []string{}
	lineGroups := [][]string{}
	for _, line := range lines {
		if idx < 3 {
			lineGroup = append(lineGroup, line)
			idx += 1
		} else {
			lineGroups = append(lineGroups, lineGroup)
			idx = 0
			lineGroup = nil
		}
	}
	// append last group
	lineGroups = append(lineGroups, lineGroup)

	for _, group := range lineGroups {
		for _, ruckSack := range group {
			m := toMap(ruckSack)
		}
	}
	return len(filename)
}
