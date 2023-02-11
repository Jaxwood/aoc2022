package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ROCK = '#'
	AIR  = '.'
	SAND = 's'
)

type Coord struct {
	X int
	Y int
}

type Reservoir = map[Coord]rune

func fill(reservoir Reservoir, from Coord, to Coord) Reservoir {
	for y := from.Y; y <= to.Y; y++ {
		for x := from.X; x <= to.X; x++ {
			if _, ok := reservoir[Coord{X: x, Y: y}]; !ok {
				reservoir[Coord{x, y}] = AIR
			}
		}
	}
	return reservoir
}

func draw(reservoir Reservoir, from Coord, to Coord) {
	for y := from.Y; y <= to.Y; y++ {
		var line string
		for x := from.X; x <= to.X; x++ {
			if val, ok := reservoir[Coord{X: x, Y: y}]; ok {
				line += string(val)
			} else {
				line += string(AIR)
			}
		}
		fmt.Println(line)
	}
}

func until(raw string) Coord {
	nums := strings.Split(raw, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return Coord{X: x, Y: y}
}

func parse(file string) Reservoir {
	lines := strings.Split(file, "\n")
	result := Reservoir{}
	for _, line := range lines {
		segments := strings.Split(line, " -> ")
		from := until(segments[0])
		rest := segments[1:]
		for _, segment := range rest {
			to := until(segment)
			xMin := math.Min(float64(from.X), float64(to.X))
			xMax := math.Max(float64(from.X), float64(to.X))
			yMin := math.Min(float64(from.Y), float64(to.Y))
			yMax := math.Max(float64(from.Y), float64(to.Y))
			for x := xMin; x <= xMax; x++ {
				for y := yMin; y <= yMax; y++ {
					result[Coord{X: int(x), Y: int(y)}] = ROCK
				}
			}
			from = to
		}
	}
	return result
}

func isResting(reservoir *Reservoir, coord Coord) bool {
	_, okLeft := (*reservoir)[Coord{coord.X - 1, coord.Y + 1}]
	_, okRight := (*reservoir)[Coord{coord.X + 1, coord.Y + 1}]

	if !okLeft || !okRight {
		return false
	}
	return true
}

func drop(reservoir *Reservoir, start Coord) (bool, Coord) {
	if next, ok := moveDown(reservoir, start); ok {
		return drop(reservoir, next)
	} else if next, ok := moveDownLeft(reservoir, start); ok {
		return drop(reservoir, next)
	} else if next, ok := moveDownRight(reservoir, start); ok {
		return drop(reservoir, next)
	} else if isResting(reservoir, start) {
		return true, start
	}
	// into the void
	return false, start
}

func moveDown(reservoir *Reservoir, coord Coord) (Coord, bool) {
	next := Coord{coord.X, coord.Y + 1}
	down, ok := (*reservoir)[next]
	if ok && down == AIR {
		return next, ok
	}
	return coord, false
}

func moveDownLeft(reservoir *Reservoir, coord Coord) (Coord, bool) {
	downLeft := Coord{coord.X - 1, coord.Y + 1}
	leftDownType, downLeftOK := (*reservoir)[downLeft]
	if downLeftOK && leftDownType == AIR {
		return downLeft, true
	}
	return coord, false
}

func moveDownRight(reservoir *Reservoir, coord Coord) (Coord, bool) {
	downRight := Coord{coord.X + 1, coord.Y + 1}
	rightDownType, downRightOK := (*reservoir)[downRight]
	if downRightOK && rightDownType == AIR {
		return downRight, true
	}
	return coord, false
}

func countSand(reservoir *Reservoir) int {
	total := 0
	for _, y := range *reservoir {
		if y == SAND {
			total += 1
		}
	}
	return total
}

func boundaries(reservoir Reservoir) (Coord, Coord) {
	xMin, xMax := 500.0, 500.0
	yMin, yMax := 0.0, 0.0
	for coord := range reservoir {
		xMin = math.Min(float64(coord.X), xMin)
		xMax = math.Max(float64(coord.X), xMax)
		yMin = math.Min(float64(coord.Y), yMin)
		yMax = math.Max(float64(coord.Y), yMax)
	}
	return Coord{int(xMin), int(yMin)}, Coord{int(xMax), int(yMax)}
}

func day14(file string) int {
	reservoir := parse(file)
	from, to := boundaries(reservoir)
	reservoir = fill(reservoir, from, to)

	for {
		resting, coord := drop(&reservoir, Coord{500, 0})
		if !resting {
			break
		} else {
			reservoir[coord] = SAND
		}
		//draw(reservoir, from, to)
	}

	return countSand(&reservoir)
}
