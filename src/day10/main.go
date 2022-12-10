package main

import (
	"strconv"
	"strings"
)

type Register = int

type Instruction interface {
	Start(int) Instruction
	Complete(int) bool
	Result(Register) Register
}

type Noop struct {
	tick int
}

type Add struct {
	tick   int
	amount int
}

func (i Add) Start(tick int) Instruction {
	i.tick = tick + 2
	return i
}

func (i Add) Complete(tick int) bool {
	return i.tick == tick
}

func (i Add) Result(r Register) Register {
	return r + i.amount
}

func (i Noop) Start(tick int) Instruction {
	i.tick = tick + 1
	return i
}

func (i Noop) Complete(tick int) bool {
	return i.tick == tick
}

func (i Noop) Result(r Register) Register {
	return r
}

func parse(file string) []Instruction {
	result := []Instruction{}
	lines := strings.Split(file, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		segments := strings.Split(line, " ")
		if len(segments) == 2 {
			amount, _ := strconv.Atoi(segments[1])
			result = append(result, Add{tick: 0, amount: amount})
		} else {
			result = append(result, Noop{0})
		}
	}
	return result
}

func day10(file string) int {
	instructions := parse(file)
	next := instructions[0]
	instructions = instructions[1:]
	next = next.Start(1)
	cycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	result := 1
	total := 0
	for cycle := 1; cycle <= 220; cycle++ {
		if next.Complete(cycle) {
			result = next.Result(result)
			next = instructions[0]
			instructions = instructions[1:]
			next = next.Start(cycle)
		}
		if _, ok := cycles[cycle]; ok {
			total += result * cycle
		}
	}

	return total
}
