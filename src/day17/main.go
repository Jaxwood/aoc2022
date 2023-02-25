package main

import (
	"fmt"
	"math"
)

type Coord struct {
	x int
	y int
}

type Shape struct {
	height int
	coords map[Coord]bool
}

var shapes = []Shape{
	{
		1,
		map[Coord]bool{
			{0, 0}: true,
			{1, 0}: true,
			{2, 0}: true,
			{3, 0}: true,
		},
	},
	{
		3,
		map[Coord]bool{
			{1, 0}: true,
			{1, 1}: true,
			{1, 2}: true,
			{0, 1}: true,
			{2, 1}: true,
		},
	},
	{
		3,
		map[Coord]bool{
			{0, 2}: true,
			{1, 2}: true,
			{2, 2}: true,
			{2, 0}: true,
			{2, 1}: true,
		},
	},
	{
		4,
		map[Coord]bool{
			{0, 0}: true,
			{0, 1}: true,
			{0, 2}: true,
			{0, 3}: true,
		},
	},
	{
		2,
		map[Coord]bool{
			{0, 0}: true,
			{0, 1}: true,
			{1, 0}: true,
			{1, 1}: true,
		},
	},
}

// move the shape in the direction of the gas
// gas is one of the following: E, W
func (s *Shape) Move(gas rune, chamber *map[Coord]bool) bool {
	return false
}

func (s *Shape) StartPosition(height int) {
	result := map[Coord]bool{}
	for coord := range s.coords {
		result[Coord{coord.x + 2, coord.y - height - 3 - s.height}] = true
	}

	s.coords = result
}

func (s *Shape) Draw() {
	for y := 0; y < 4; y++ {
		line := ""
		for x := 0; x < 4; x++ {
			if _, ok := s.coords[Coord{x, y}]; ok {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func findFloorHeight(chamber map[Coord]bool) int {
	y := 0.0
	for k, _ := range chamber {
		y = math.Min(float64(k.y), y)
	}

	return int(y)
}

 func draw(chamber map[Coord]bool, shape Shape) {
	for y := -10; y <= 0; y++ {
		line := ""
		for x := 0; x < 7; x++ {
			if _, ok := chamber[Coord{x, y}]; ok {
				line += "#"
			} else if _, ok := shape.coords[Coord{x, y}]; ok {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
 }

func day17(input string, until int) int {
	// add the floor
	chamber := map[Coord]bool{}
	for i := 0; i < 7; i++ {
		chamber[Coord{i, 0}] = true
	}

	gasSize := len(input)
	gasIdx := 0
	shapeSize := len(shapes)
	shapeIdx := 0
	for i := 0; i < until; i++ {
		shape := shapes[shapeIdx]
		height := findFloorHeight(chamber)
		shape.StartPosition(height)
		draw(chamber, shape)
		shapeIdx = (shapeIdx + 1) % shapeSize
		for {
			gas := rune(input[gasIdx])
			gasIdx = (gasIdx + 1) % gasSize
			if !shape.Move(gas, &chamber) {
				break
			}
		}
	}

	return 0
}
