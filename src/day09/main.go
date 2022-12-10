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
	return g
}

func (r MoveLeft) Update(g Grid) Grid {
	g.Head.X -= 1
	return g
}

func (r MoveUp) Update(g Grid) Grid {
	g.Head.Y += 1
	return g
}

func (r MoveDown) Update(g Grid) Grid {
	g.Head.Y -= 1
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

func (g Grid) SameColumn() bool {
	return g.Head.Y == g.Tail.Y
}

func (g Grid) SameRow() bool {
	return g.Head.X == g.Tail.X
}

func (g Grid) Diagonal() bool {
	return !g.SameColumn() && !g.SameRow()
}

func (g Grid) OnTop() bool {
	return g.Head == g.Tail
}

func (g Grid) Adjecent() bool {
	// check neighbours
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			candidate := Coord{g.Tail.X + i, g.Tail.Y + j}
			if candidate == g.Head {
				return true
			}
		}
	}

	return false
}

func (g Grid) Move() Grid {
	// if tail is next to or on top of head - don't move it
	if g.OnTop() || g.Adjecent() {
		return g
	}
	// move diagonal
	if g.Diagonal() {
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
	// move left/right
	if g.SameColumn() {
		if g.Head.X > g.Tail.X {
			g.Tail.X += 1
		} else {
			g.Tail.X -= 1
		}
		g.Visited[g.Tail] = true
		return g
	}
	// move up/down
	if g.SameRow() {
		if g.Head.Y > g.Tail.Y {
			g.Tail.Y += 1
		} else {
			g.Tail.Y -= 1
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
		grid = grid.Move()
	}
	return len(grid.Visited)
}

func day09b(file string) int {
	moves := parse(file)
	grids := map[int]Grid{}
	for i := 0; i < 10; i++ {
		grid := Grid{
			map[Coord]bool{
				Coord{0, 0}: true,
			},
			Coord{0, 0},
			Coord{0, 0},
		}
		grids[i] = grid
	}
	for _, move := range moves {
		// move head of rope
		grid := grids[0]
		grid = move.Update(grid)
		grid = grid.Move()
		grids[0] = grid
		last := grid.Head
		// move all knots
		for i := 1; i < 10; i++ {
			grid := grids[i]
			grid.Head = last
			grid = grid.Move()
			grids[i] = grid
			last = grid.Tail
		}
	}
	return len(grids[9].Visited)
}
