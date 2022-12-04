package main

import (
	"strconv"
	"strings"

	"github.com/jaxwood/aoc2022/internal/maps"
)

func parse(lines []string) [][]map[int]bool {
	assignments := [][]map[int]bool{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		sections := strings.Split(line, ",")
		assignment := []map[int]bool{}
		for _, section := range sections {
			set := map[int]bool{}
			ranges := strings.Split(section, "-")
			start, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])
			for i := start; i <= end; i++ {
				set[i] = true
			}
			assignment = append(assignment, set)
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
		diff := maps.Difference(assignment[0], assignment[1])
		diff2 := maps.Difference(assignment[1], assignment[0])
		if len(diff) == 0 || len(diff2) == 0 {
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
		diff := maps.Difference(assignment[0], assignment[1])
		diff2 := maps.Difference(assignment[1], assignment[0])
		if len(diff) != len(assignment[0]) || len(diff2) != len(assignment[1]) {
			total += 1
		}
	}
	return total
}
