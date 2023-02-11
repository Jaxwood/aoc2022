package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Reservoir = map[Coord]rune

func draw(reservoir Reservoir) {
	for y := 0; y < 15; y++ {
		var line string
		for x := 490; x < 510; x++ {
			if val, ok := reservoir[Coord{X: x, Y: y}]; ok {
				line += string(val)
			} else {
				line += "."
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
					result[Coord{X: int(x), Y: int(y)}] = '#'
				}
			}
			from = to
		}
	}
	return result
}

func day14(file string) int {
	reservoir := parse(file)
	draw(reservoir)
	return 0
}
