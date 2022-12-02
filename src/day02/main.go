package main

import (
	"strings"
)

const ROCK = "A"
const PAPER = "B"
const SCISSOR = "C"

const ROCK_PLAYER = "X"
const PAPER_PLAYER = "Y"
const SCISSOR_PLAYER = "Z"

func score(round []string) int {
	score := 0
	hand1 := strings.TrimSpace(round[0])
	hand2 := strings.TrimSpace(round[1])
	// score the players card
	if hand2 == ROCK_PLAYER {
		score += 1
	} else if hand2 == PAPER_PLAYER {
		score += 2
	} else if hand2 == SCISSOR_PLAYER {
		score += 3
	}

	// draw
	if hand1 == ROCK && hand2 == ROCK_PLAYER {
		score += 3
	} else if hand1 == PAPER && hand2 == PAPER_PLAYER {
		score += 3
	} else if hand1 == SCISSOR && hand2 == SCISSOR_PLAYER {
		score += 3
	} else if hand1 == ROCK && hand2 == PAPER_PLAYER {
		score += 6
	} else if hand1 == PAPER && hand2 == SCISSOR_PLAYER {
		score += 6
	} else if hand1 == SCISSOR && hand2 == ROCK_PLAYER {
		score += 6
	}

	return score
}

func day02a(filename string) int {
	lines := strings.Split(filename, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		segments := strings.Split(line, " ")
		total += score(segments)
	}

	return total
}

