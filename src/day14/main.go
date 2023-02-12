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

func (r *Reservoir) draw() {
	for y := r.yMin; y <= r.yMax; y++ {
		var line string
		for x := r.xMin; x <= r.xMax; x++ {
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

func parse(file string) map[Coord]rune {
	lines := strings.Split(file, "\n")
	rocks := map[Coord]rune{}
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
					rocks[Coord{X: int(x), Y: int(y)}] = ROCK
				}
			}
			from = to
		}
	}
	return rocks
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

func (r *Reservoir) isResting(coord Coord) bool {
	left := Coord{coord.X - 1, coord.Y + 1}
	right := Coord{coord.X + 1, coord.Y + 1}

	if r.withIn(left) && r.withIn(right) {
		return true
	}
	return false
}

func (r *Reservoir) moveDown(coord Coord) (Coord, bool) {
	next := Coord{coord.X, coord.Y + 1}
	_, ok := r.Rocks[next]
	if r.withIn(next) && !ok {
		return next, true
	}
	return coord, false
}

func (r *Reservoir) moveDownLeft(coord Coord) (Coord, bool) {
	next := Coord{coord.X - 1, coord.Y + 1}
	_, ok := r.Rocks[next]
	if r.withIn(next) && !ok {
		return next, true
	}
	return coord, false
}

func (r *Reservoir) moveDownRight(coord Coord) (Coord, bool) {
	next := Coord{coord.X + 1, coord.Y + 1}
	_, ok := r.Rocks[next]
	if r.withIn(next) && !ok {
		return next, true
	}
	return coord, false
}

func (r *Reservoir) withIn(coord Coord) bool {
	if coord.X < r.xMin || coord.X > r.xMax {
		return false
	}
	if coord.Y < r.yMin || coord.Y > r.yMax {
		return false
	}
	return true
}

func (r *Reservoir) count() int {
	total := 0
	for _, y := range r.Rocks {
		if y == SAND {
			total += 1
		}
	}
	return total
}

func boundaries(rocks map[Coord]rune) (Coord, Coord) {
	xMin, xMax := 500.0, 500.0
	yMin, yMax := 0.0, 0.0
	for coord := range rocks {
		xMin = math.Min(float64(coord.X), xMin)
		xMax = math.Max(float64(coord.X), xMax)
		yMin = math.Min(float64(coord.Y), yMin)
		yMax = math.Max(float64(coord.Y), yMax)
	}
	return Coord{int(xMin), int(yMin)}, Coord{int(xMax), int(yMax)}
}

func day14(file string) int {
	rocks := parse(file)
	min, max := boundaries(rocks)
	reservoir := Reservoir{
		rocks,
		min.X,
		max.X,
		min.Y,
		max.Y,
	}

	for {
		resting, coord := reservoir.drop(Coord{500, 0})
		if !resting {
			break
		}
		reservoir.Rocks[coord] = SAND
		//reservoir.draw()
	}

	return reservoir.count()
}

func day14b(file string) int {
	rocks := parse(file)
	min, max := boundaries(rocks)
	// add floor
	floor := 1000
	for x := min.X - floor; x <= max.X + floor; x++ {
		rocks[Coord{x, max.Y + 2}] = ROCK
	}
	reservoir := Reservoir{
		rocks,
		min.X - floor,
		max.X + floor,
		min.Y,
		max.Y + 2,
	}

	start := Coord{500, 0}
	for {
		_, coord := reservoir.drop(start)
		if coord ==  start {
			reservoir.Rocks[coord] = SAND
			break
		}
		reservoir.Rocks[coord] = SAND
		// reservoir.draw()
	}

	return reservoir.count() 
}
