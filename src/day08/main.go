package main

import (
	"math"
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

func (f Forest) left(coord Pair) (bool, int) {
	cnt := 0
	for i := coord.X - 1; i >= 0; i-- {
		val := f.trees[Pair{i, coord.Y}]
		if val >= f.trees[coord] {
			return false, cnt + 1
		}
		cnt += 1
	}
	return true, cnt
}

func (f Forest) right(coord Pair) (bool, int) {
	cnt := 0
	for i := coord.X + 1; i < f.width; i++ {
		val := f.trees[Pair{i, coord.Y}]
		if val >= f.trees[coord] {
			return false, cnt + 1
		}
		cnt += 1
	}
	return true, cnt
}

func (f Forest) up(coord Pair) (bool, int) {
	cnt := 0
	for i := coord.Y - 1; i >= 0; i-- {
		val := f.trees[Pair{coord.X, i}]
		if val >= f.trees[coord] {
			return false, cnt + 1
		}
		cnt += 1
	}
	return true, cnt
}

func (f Forest) down(coord Pair) (bool, int) {
	cnt := 0
	for i := coord.Y + 1; i < f.height; i++ {
		val := f.trees[Pair{coord.X, i}]
		if val >= f.trees[coord] {
			return false, cnt + 1
		}
		cnt += 1
	}
	return true, cnt
}

func (f Forest) edge(coord Pair) bool {
	if coord.X == 0 || coord.Y == 0 || coord.X == f.width-1 || coord.Y == f.height-1 {
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
		} else if ok, _ := forest.left(k); ok {
			total += 1
		} else if ok, _ = forest.right(k); ok {
			total += 1
		} else if ok, _ := forest.up(k); ok {
			total += 1
		} else if ok, _ := forest.down(k); ok {
			total += 1
		}
	}
	return total
}

func day08b(file string) float64 {
	best := 0.0
	forest := parse(file)
	for k, _ := range forest.trees {
		_, left := forest.left(k)
		_, right := forest.right(k)
		_, up := forest.up(k)
		_, down := forest.down(k)
		best = math.Max(best, float64(left*right*up*down))
	}
	return best
}
