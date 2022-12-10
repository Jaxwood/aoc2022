package main

import (
	"strconv"
	"strings"
)

type Move interface {
	Update(Grid) Grid
}

type MoveRight struct {
}

type MoveLeft struct {
}

type MoveUp struct {
}

type MoveDown struct {
}

func (r MoveRight) Update(g Grid) Grid {
	g.Head.X += 1
	g = g.Move()
	return g
}

func (r MoveLeft) Update(g Grid) Grid {
	g.Head.X -= 1
	g = g.Move()
	return g
}

func (r MoveUp) Update(g Grid) Grid {
	g.Head.Y += 1
	g = g.Move()
	return g
}

func (r MoveDown) Update(g Grid) Grid {
	g.Head.Y -= 1
	g = g.Move()
	return g
}

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Visited map[Coord]bool
	Head    Coord
	Tail    Coord
}

func (g Grid) AdjecentToCoord(coord Coord) []Coord {
	result := []Coord{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			candidate := Coord{coord.X + i, coord.Y + j}
			result = append(result, candidate)
		}
	}
	return result
}

func (g Grid) SameColumn() bool {
	return g.Head.Y == g.Tail.Y
}

func (g Grid) SameRow() bool {
	return g.Head.X == g.Tail.X
}

func (g Grid) AdjecentOrOnTop() bool {
	// if tail is ontop of head - don't move it
	if g.Head == g.Tail {
		return true
	}
	// check neighbours
	adjecent := g.AdjecentToCoord(g.Tail)
	for _, coord := range adjecent {
		if coord == g.Head {
			return true
		}
	}

	return false
}

func (g Grid) Move() Grid {
	// if tail is next to or on top of head - don't move it
	if g.AdjecentOrOnTop() {
		return g
	}
	// move diagonal
	if !g.SameColumn() && !g.SameRow() {
		if g.Head.X > g.Tail.X && g.Head.Y > g.Tail.Y {
			g.Tail.X += 1
			g.Tail.Y += 1
		} else if g.Head.X > g.Tail.X && g.Head.Y < g.Tail.Y {
			g.Tail.X += 1
			g.Tail.Y -= 1
		} else if g.Head.X < g.Tail.X && g.Head.Y > g.Tail.Y {
			g.Tail.X -= 1
			g.Tail.Y += 1
		} else if g.Head.X < g.Tail.X && g.Head.Y < g.Tail.Y {
			g.Tail.X -= 1
			g.Tail.Y -= 1
		}
		g.Visited[g.Tail] = true
		return g
	}
	// move up/down
	if !g.SameColumn() {
		if g.Head.Y > g.Tail.Y {
			g.Tail.Y += 1
		} else {
			g.Tail.Y -= 1
		}
		g.Visited[g.Tail] = true
		return g
	}
	// move left/right
	if !g.SameRow() {
		if g.Head.X > g.Tail.X {
			g.Tail.X += 1
		} else {
			g.Tail.X -= 1
		}
		g.Visited[g.Tail] = true
		return g
	}
	return g
}

func parse(file string) []Move {
	result := []Move{}
	lines := strings.Split(file, "\n")
	for _, line := range lines {
		segments := strings.Split(line, " ")
		if len(segments) == 2 {
			amount, _ := strconv.Atoi(segments[1])
			for i := 0; i < amount; i++ {
				switch segments[0] {
				case "R":
					result = append(result, MoveRight{})
				case "L":
					result = append(result, MoveLeft{})
				case "U":
					result = append(result, MoveUp{})
				case "D":
					result = append(result, MoveDown{})
				}
			}
		}
	}
	return result
}

func day09(file string) int {
	moves := parse(file)
	grid := Grid{
		map[Coord]bool{
			Coord{0, 0}: true,
		},
		Coord{0, 0},
		Coord{0, 0},
	}
	for _, move := range moves {
		grid = move.Update(grid)
	}
	return len(grid.Visited)
}
