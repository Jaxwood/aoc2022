package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Cargo struct {
	Items map[int][]rune
}

type Move struct {
	Count int
	From  int
	To    int
}

func (m Move) String() string {
	return fmt.Sprintf("move %v from %v to %v", m.Count, m.From, m.To)
}

func (c Cargo) String() string {
	i := 1
	result := ""
	for true {
		if val, ok := c.Items[i]; ok {
			result += string(val[len(val)-1])
		} else {
			break
		}
		i += 1
	}
	return result
}

func (c Cargo) Move(move Move) {
	for i := 0; i < move.Count; i++ {
		slice := c.Items[move.From]
		if len(slice) > 0 {
			item := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			c.Items[move.To] = append(c.Items[move.To], item)
			c.Items[move.From] = slice
		}
	}
}

func (c Cargo) MoveBlock(move Move) {
	block := []rune{}
	for i := 0; i < move.Count; i++ {
		slice := c.Items[move.From]
		if len(slice) > 0 {
			item := slice[len(slice)-1]
			block = append(block, item)
			slice = slice[:len(slice)-1]
			c.Items[move.From] = slice
		}
	}
	for i := len(block) - 1; i >= 0; i-- {
		c.Items[move.To] = append(c.Items[move.To], block[i])
	}
}

func parse(filename string) []Move {
	lines := strings.Split(filename, "\n")
	moves := []Move{}
	for _, line := range lines {
		if strings.HasPrefix(line, "move") {
			segments := strings.Split(line, " ")
			count, _ := strconv.Atoi(segments[1])
			from, _ := strconv.Atoi(segments[3])
			to, _ := strconv.Atoi(segments[5])
			moves = append(moves, Move{count, from, to})
		}
	}
	return moves
}

func day05a(filename string, cargo Cargo) string {
	moves := parse(filename)
	for _, move := range moves {
		cargo.Move(move)
	}
	return cargo.String()
}

func day05b(filename string, cargo Cargo) string {
	moves := parse(filename)
	for _, move := range moves {
		cargo.MoveBlock(move)
	}
	return cargo.String()
}
