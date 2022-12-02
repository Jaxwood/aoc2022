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

const LOOSE = "X"
const DRAW = "Y"
const WIN = "Z"

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

func cardToPlay(card string, outcome string) string {
	if card == ROCK && outcome == DRAW {
		return ROCK_PLAYER
	}
	if card == ROCK && outcome == LOOSE {
		return SCISSOR_PLAYER
	}
	if card == ROCK && outcome == WIN {
		return PAPER_PLAYER
	}
	if card == PAPER && outcome == DRAW {
		return PAPER_PLAYER
	}
	if card == PAPER && outcome == LOOSE {
		return ROCK_PLAYER
	}
	if card == PAPER && outcome == WIN {
		return SCISSOR_PLAYER
	}
	if card == SCISSOR && outcome == DRAW {
		return SCISSOR_PLAYER
	}
	if card == SCISSOR && outcome == LOOSE {
		return PAPER_PLAYER
	}
	if card == SCISSOR && outcome == WIN {
		return ROCK_PLAYER
	}
	panic("unknown case")
}

func scoreb(round []string) int {
	total := 0
	opponent := strings.TrimSpace(round[0])
	player := strings.TrimSpace(round[1])

	card := cardToPlay(opponent, player)

	return total + score([]string{ opponent, card })
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

func day02b(filename string) int {
	lines := strings.Split(filename, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		segments := strings.Split(line, " ")
		total += scoreb(segments)
	}

	return total
}

