package main

import (
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Forest struct {
	trees  map[Pair]int
	width  int
	height int
}

func (f Forest) left(coord Pair) bool {
	for i := 0; i < coord.X; i++ {
		val := f.trees[Pair{i, coord.Y}]
		if val >= f.trees[coord] {
			return false
		}
	}
	return true
}

func (f Forest) right(coord Pair) bool {
	for i := coord.X + 1; i < f.width; i++ {
		val := f.trees[Pair{i, coord.Y}]
		if val >= f.trees[coord] {
			return false
		}
	}
	return true
}

func (f Forest) up(coord Pair) bool {
	for i := 0; i < coord.Y; i++ {
		val := f.trees[Pair{coord.X, i}]
		if val >= f.trees[coord] {
			return false
		}
	}
	return true
}

func (f Forest) down(coord Pair) bool {
	for i := coord.Y + 1; i < f.height; i++ {
		val := f.trees[Pair{coord.X, i}]
		if val >= f.trees[coord] {
			return false
		}
	}
	return true
}

func (f Forest) edge(coord Pair) bool {
	if coord.X == 0 || coord.Y == 0 || coord.X == f.width - 1 || coord.Y == f.height - 1 {
		return true
	}
	return false
}

func parse(file string) Forest {
	lines := strings.Split(file, "\n")
	result := map[Pair]int{}
	for y, line := range lines {
		for x, tree := range line {
			val, _ := strconv.Atoi(string(tree))
			result[Pair{x, y}] = val
		}
	}
	return Forest{result, len(lines[0]), len(lines) - 1}
}

func day08(file string) int {
	total := 0
	forest := parse(file)
	for k, _ := range forest.trees {
		if forest.edge(k) {
			total += 1
		} else if forest.left(k) {
			total += 1
		} else if forest.right(k) {
			total += 1
		} else if forest.up(k) {
			total += 1
		} else if forest.down(k) {
			total += 1
		}
	}
	return total
}
