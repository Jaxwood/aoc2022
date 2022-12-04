package main

import (
	"strconv"
	"strings"

	"github.com/jaxwood/aoc2022/internal/set"
)

func parse(lines []string) [][]set.Set {
	assignments := [][]set.Set{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		sections := strings.Split(line, ",")
		assignment := []set.Set{}
		for _, section := range sections {
			items := map[rune]bool{}
			ranges := strings.Split(section, "-")
			start, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])

			for i := start; i <= end; i++ {
				items[rune(i)] = true
			}
			assignment = append(assignment, set.Set{items})
		}
		assignments = append(assignments, assignment)
	}

	return assignments
}

func day04a(filename string) int {
	lines := strings.Split(filename, "\n")
	assignments := parse(lines)
	total := 0
	for _, assignment := range assignments {
		diff := assignment[0].Difference(assignment[1])
		diff2 := assignment[1].Difference(assignment[0])
		if diff.Len() == 0 || diff2.Len() == 0 {
			total += 1
		}
	}

	return total
}

func day04b(filename string) int {
	lines := strings.Split(filename, "\n")
	assignments := parse(lines)
	total := 0
	for _, assignment := range assignments {
		diff := assignment[0].Difference(assignment[1])
		diff2 := assignment[1].Difference(assignment[0])
		if diff.Len() != assignment[0].Len() || diff2.Len() != assignment[1].Len() {
			total += 1
		}
	}
	return total
}
