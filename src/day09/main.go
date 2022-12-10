package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Move interface {
	Update(Grid) Grid
}

type MoveRight struct {
	Amount int
}

type MoveLeft struct {
	Amount int
}

type MoveUp struct {
	Amount int
}

type MoveDown struct {
	Amount int
}

func (r MoveRight) Update(g Grid) Grid {
	head := g.Head
	for i := 0; i < r.Amount; i++ {
		head.X += 1
		g = g.Move(head)
	}
	return g
}

func (r MoveLeft) Update(g Grid) Grid {
	head := g.Head
	for i := 0; i < r.Amount; i++ {
		head.X -= 1
		g = g.Move(head)
	}
	return g
}

func (r MoveUp) Update(g Grid) Grid {
	head := g.Head
	for i := 0; i < r.Amount; i++ {
		head.Y += 1
		g = g.Move(head)
	}
	return g
}

func (r MoveDown) Update(g Grid) Grid {
	head := g.Head
	for i := 0; i < r.Amount; i++ {
		head.Y -= 1
		g = g.Move(head)
	}
	return g
}

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Visited map[Coord]bool
	Head Coord
	Tail Coord
}

func (g Grid) Move(head Coord) Grid {
	g.Head = head
	xDiff := g.Head.X - g.Tail.X
	yDiff := g.Head.Y - g.Tail.Y

	return g
}

func parse(file string) []Move {
	result := []Move{}
	lines := strings.Split(file, "\n")
	for _, line := range lines {
		segments := strings.Split(line, " ")
		if len(segments) == 2 {
			amount, _ := strconv.Atoi(segments[1])
			switch segments[0] {
				case "R":
				result = append(result, MoveRight{amount})
				case "L":
				result = append(result, MoveLeft{amount})
				case "U":
				result = append(result, MoveUp{amount})
				case "D":
				result = append(result, MoveDown{amount})
			}

		}
	}
	return result
}

func day09(file string) int {
	moves := parse(file)
	grid := Grid{
		map[Coord]bool{
			Coord{0,0}: true,
		},
		Coord{0,0},
		Coord{0,0},
	}
	for _, move := range moves {
		grid = move.Update(grid)
	}
	fmt.Println(grid.Head)
	return len(grid.Visited)
}
