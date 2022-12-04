package main

import (
	"strings"

	"github.com/jaxwood/aoc2022/internal/maps"
)

const (
	a = 97
	z = 123
	A = 65
	Z = 92
)

func scoreTable() map[rune]int {
	scores := map[rune]int{}

	// lowercases
	for i := a; i < z; i++ {
		scores[rune(i)] = i - 96
	}
	// uppercases
	for i := A; i < Z; i++ {
		scores[rune(i)] = 27 + (i - A)
	}

	return scores
}

func day03a(filename string) int {
	lines := strings.Split(filename, "\n")
	sum := 0
	scores := scoreTable()
	for _, line := range lines {
		start := line[:len(line)/2]
		end := line[len(line)/2:]
		startChars := maps.StrToMap(start)
		for _, c := range end {
			_, ok := startChars[c]
			if ok {
				sum += scores[c]
				break
			}
		}
	}
	return sum
}

func ruckSackGroups(lines []string) [][]maps.Set {
	idx := 0
	ruckSack := []maps.Set{}
	ruckSacks := [][]maps.Set{}

	for _, line := range lines {
		if idx < 3 {
			ruckSack = append(ruckSack, maps.Set{maps.StrToMap(line)})
			idx += 1
		} else {
			ruckSacks = append(ruckSacks, ruckSack)
			idx = 1
			ruckSack = nil
			ruckSack = append(ruckSack, maps.Set{maps.StrToMap(line)})
		}
	}
	// append last group
	ruckSacks = append(ruckSacks, ruckSack)

	return ruckSacks
}

func day03b(filename string) int {
	sum := 0
	scores := scoreTable()
	lines := strings.Split(filename, "\n")
	// group by lines of 3
	ruckSacks := ruckSackGroups(lines)

	for _, group := range ruckSacks {
		first := group[0]
		rest := group[1:]
		for _, ruckSack := range rest {
			first = first.Intersect(ruckSack)
		}
		for k, _ := range first.Items {
			sum += scores[k]
		}
	}

	return sum
}
