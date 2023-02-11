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

type Reservoir struct {
	Rocks map[Coord]rune
	xMin  int
	xMax  int
	yMin  int
	yMax  int
}

func (r *Reservoir) fill() {
	for y := r.yMin; y <= r.yMax; y++ {
		for x := r.xMin; x <= r.xMax; x++ {
			if _, ok := r.Rocks[Coord{X: x, Y: y}]; !ok {
				r.Rocks[Coord{x, y}] = AIR
			}
		}
	}
}

func (r *Reservoir) draw(from Coord, to Coord) {
	for y := from.Y; y <= to.Y; y++ {
		var line string
		for x := from.X; x <= to.X; x++ {
			if val, ok := r.Rocks[Coord{X: x, Y: y}]; ok {
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
	result := Reservoir{
		map[Coord]rune{},
		0,
		0,
		0,
		0,
	}
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
					result.Rocks[Coord{X: int(x), Y: int(y)}] = ROCK
				}
			}
			from = to
		}
	}
	result.boundaries()
	return result
}

func (r *Reservoir) isResting(coord Coord) bool {
	_, okLeft := r.Rocks[Coord{coord.X - 1, coord.Y + 1}]
	_, okRight := r.Rocks[Coord{coord.X + 1, coord.Y + 1}]

	if !okLeft || !okRight {
		return false
	}
	return true
}

func (r *Reservoir) drop(start Coord) (bool, Coord) {
	if next, ok := r.moveDown(start); ok {
		return r.drop(next)
	} else if next, ok := r.moveDownLeft(start); ok {
		return r.drop(next)
	} else if next, ok := r.moveDownRight(start); ok {
		return r.drop(next)
	} else if r.isResting(start) {
		return true, start
	}
	// into the void
	return false, start
}

func (r *Reservoir) moveDown(coord Coord) (Coord, bool) {
	next := Coord{coord.X, coord.Y + 1}
	down, ok := r.Rocks[next]
	if ok && down == AIR {
		return next, ok
	}
	return coord, false
}

func (r *Reservoir) moveDownLeft(coord Coord) (Coord, bool) {
	downLeft := Coord{coord.X - 1, coord.Y + 1}
	leftDownType, downLeftOK := r.Rocks[downLeft]
	if downLeftOK && leftDownType == AIR {
		return downLeft, true
	}
	return coord, false
}

func (r *Reservoir) moveDownRight(coord Coord) (Coord, bool) {
	downRight := Coord{coord.X + 1, coord.Y + 1}
	rightDownType, downRightOK := r.Rocks[downRight]
	if downRightOK && rightDownType == AIR {
		return downRight, true
	}
	return coord, false
}

func (r *Reservoir) countSand() int {
	total := 0
	for _, y := range r.Rocks {
		if y == SAND {
			total += 1
		}
	}
	return total
}

func (r *Reservoir) boundaries() {
	xMin, xMax := 500.0, 500.0
	yMin, yMax := 0.0, 0.0
	for coord := range r.Rocks {
		xMin = math.Min(float64(coord.X), xMin)
		xMax = math.Max(float64(coord.X), xMax)
		yMin = math.Min(float64(coord.Y), yMin)
		yMax = math.Max(float64(coord.Y), yMax)
	}
	r.xMin = int(xMin)
	r.xMax = int(xMax)
	r.yMin = int(yMin)
	r.yMax = int(yMax)
}

func day14(file string) int {
	reservoir := parse(file)
	reservoir.fill()

	for {
		resting, coord := reservoir.drop(Coord{500, 0})
		if !resting {
			break
		} else {
			reservoir.Rocks[coord] = SAND
		}
		//draw(reservoir, from, to)
	}

	return reservoir.countSand()
}

func day14b(file string) int {
	return len(file)
}
