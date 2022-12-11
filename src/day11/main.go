package main

import (
	"strings"
	"strconv"
)

type Operation int

const (
	Divide Operation = 0
	Multiply = 1
	Add = 2
)

type Monkey struct {
	Num int
	Levels []int
	Operation Operation
	Operand int
	Test int
	Truecase int
	Falsecase int
}

func parse(file string) map[int]Monkey {
	result := map[int]Monkey{}
	lines := strings.Split(file, "\n")
	// levels := regexp.MustCompile("Starting items: (\\d+)")
	var monkey Monkey
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Monkey") {
			segments := strings.Split(line, " ")
			numStr := segments[1]
			num, _ := strconv.Atoi(numStr[:1])
			monkey = Monkey{}
			monkey.Num = num
		}
		if strings.Contains(line, "Starting") {
			levels := strings.Split(line, " ")[4:]
			for _, level := range levels {
				l, _ := strconv.Atoi(strings.TrimSuffix(level, ","))
				monkey.Levels = append(monkey.Levels, l)
			}
		}
		if strings.Contains(line, "Operation") {
			operations := strings.Split(line, " ")[6:]
			switch operations[0] {
			case "*":
				monkey.Operation = Multiply
			case "+":
				monkey.Operation = Add
			}
			num, _ := strconv.Atoi(operations[1])
			// if num = 0 - apply operation to self
			monkey.Operand = num
		}
		if strings.Contains(line, "Test:") {
			tests := strings.Split(line, " ")[5:]
			num, _ := strconv.Atoi(tests[0])
			monkey.Test = num
		}
		if strings.Contains(line, "true:") {
			target := strings.Split(line, " ")[9:]
			num, _ := strconv.Atoi(target[0])
			monkey.Truecase = num
		}
		if strings.Contains(line, "false:") {
			target := strings.Split(line, " ")[9:]
			num, _ := strconv.Atoi(target[0])
			monkey.Falsecase = num
			result[monkey.Num] = monkey
		}
	}
	return result
}

func day11a(file string) int {
	monkeys := parse(file)
	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			queue := monkey.Levels
			for len(queue) > 0 {
				next := queue[0]
				queue = queue[1:]
				// calculate new value
				switch monkey.Operation {
				case 1:
					if monkey.Operand == 0 {
						next *= next
					} else {
						next *= monkey.Operand
					}
				case 2:
					if monkey.Operand == 0 {
						next += next
					} else {
						next += monkey.Operand
					}
				}
				// throw to other monkey
				if next % monkey.Test == 0 {
					m := monkeys[monkey.Truecase]
					m.Levels = append(m.Levels, next)
					monkeys[monkey.Truecase] = m
				} else {
					m := monkeys[monkey.Falsecase]
					m.Levels = append(m.Levels, next)
					monkeys[monkey.Falsecase] = m
				}
			}
		}
	}
	return 0
}
