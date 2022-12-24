package main

import (
	"math"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type HeightMap = map[Coord]int

func parse(file string) HeightMap {
	lines := strings.Split(file, "\n")
	result := HeightMap{}
	for y, line := range lines {
		for x, alt := range line {
			result[Coord{x, y}] = int(alt)
		}
	}
	return result
}

func findLocation(heightMap HeightMap, location rune) Coord {
	for k, v := range heightMap {
		if v == int(location) {
			return k
		}
	}
	panic("could not locate start location")
}

func neighbours(candidate Coord) []Coord {
	x := candidate.X
	y := candidate.Y
	return []Coord{
		Coord{x - 1, y},
		Coord{x + 1, y},
		Coord{x, y - 1},
		Coord{x, y + 1},
	}
}

func valid(heightMap HeightMap, from Coord, to Coord) bool {
	return heightMap[to]-1 <= heightMap[from]
}

func distance(heightMap HeightMap, start Coord, end Coord) int {
	unvisited := map[Coord]int{}
	for k, _ := range heightMap {
		unvisited[k] = math.MaxInt32
	}
	unvisited[start] = 0
	queue := []Coord{start}
	for len(queue) > 0 {
		next := queue[0]
		if next == end {
			return unvisited[end]
		}
		queue = queue[1:]
		move := unvisited[next]
		delete(unvisited, next)
		locations := neighbours(next)
		for _, loc := range locations {
			if val, ok := unvisited[loc]; ok {
				if move+1 < val && valid(heightMap, next, loc) {
					unvisited[loc] = move + 1
					queue = append(queue, loc)
				}
			}
		}
	}
	return 0
}

func day12(file string) int {
	heightMap := parse(file)
	start := findLocation(heightMap, 'S')
	end := findLocation(heightMap, 'E')
	heightMap[start] = int('a')
	heightMap[end] = int('z')

	return distance(heightMap, start, end)
}

func day12b(file string) int {
	heightMap := parse(file)
	start := findLocation(heightMap, 'S')
	end := findLocation(heightMap, 'E')
	heightMap[start] = int('a')
	heightMap[end] = int('z')

	candidates := []Coord{}
	for k, v := range heightMap {
		if v == int('a') {
			candidates = append(candidates, k)
		}
	}

	best := math.MaxInt

	for _, candidate := range candidates {
		distance := distance(heightMap, candidate, end)
		// no path to destination
		if distance == 0 {
			continue
		}
		if distance < best {
			best = distance
		}
	}

	return best
}
