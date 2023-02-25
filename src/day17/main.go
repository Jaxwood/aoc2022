package main

import "fmt"

type Coord struct {
	x int
	y int
}

type Shape struct {
	coords map[Coord]bool
}

var shapes = []Shape{
	{
		map[Coord]bool{
			{0, 0}: true,
			{1, 0}: true,
			{2, 0}: true,
			{3, 0}: true,
		},
	},
	{
		map[Coord]bool{
			{1, 0}: true,
			{1, 1}: true,
			{1, 2}: true,
			{0, 1}: true,
			{2, 1}: true,
		},
	},
	{
		map[Coord]bool{
			{0, 2}: true,
			{1, 2}: true,
			{2, 2}: true,
			{2, 0}: true,
			{2, 1}: true,
		},
	},
	{
		map[Coord]bool{
			{0, 0}: true,
			{0, 1}: true,
			{0, 2}: true,
			{0, 3}: true,
		},
	},
	{
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
func (s *Shape) Move(gas rune) {
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

func day17(input string) int {
	// add the floor
	chamber := map[Coord]bool{}
	for i := 0; i < 7; i++ {
		chamber[Coord{i, 0}] = true
	}

	size := len(input)
	idx := 0
	for _, shape := range shapes {
		shape.Draw()
		gas := rune(input[idx])
		fmt.Println("-----" + string(gas) + "-----")
		shape.Move(gas)
		idx = (idx + 1) % size
	}

	return 0
}
