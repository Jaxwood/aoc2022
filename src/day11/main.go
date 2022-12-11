package main

import (
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Operation int

const (
	Divide   Operation = 0
	Multiply           = 1
	Add                = 2
)

type Monkey struct {
	Num       int
	Levels    []*big.Int
	Operation Operation
	Operand   int
	Test      int
	Truecase  int
	Falsecase int
	Inspect   int64
}

type ByInspect []Monkey

func (a ByInspect) Len() int           { return len(a) }
func (a ByInspect) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInspect) Less(i, j int) bool { return a[i].Inspect < a[j].Inspect }

func parse(file string) map[int]Monkey {
	result := map[int]Monkey{}
	lines := strings.Split(file, "\n")
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
				monkey.Levels = append(monkey.Levels, big.NewInt(int64(l)))
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

func day11a(file string, rounds int) int64 {
	monkeys := parse(file)
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			monkey.Inspect += int64(len(monkey.Levels))
			queue := monkey.Levels
			for len(queue) > 0 {
				next := queue[0]
				queue = queue[1:]
				// inspect
				switch monkey.Operation {
				case 1:
					if monkey.Operand == 0 {
						next = big.NewInt(0).Mul(next, next)
					} else {
						next = big.NewInt(0).Mul(next, big.NewInt(int64(monkey.Operand)))
					}
				case 2:
					if monkey.Operand == 0 {
						next = big.NewInt(0).Add(next, next)
					} else {
						next = big.NewInt(0).Add(next, big.NewInt(int64(monkey.Operand)))
					}
				}
				// bored
				if rounds == 20 {
					next = big.NewInt(0).Div(next, big.NewInt(int64(3)))
				}
				// throw
				if big.NewInt(0).Mod(next, big.NewInt(int64(monkey.Test))).Cmp(big.NewInt(0)) == 0 {
					m := monkeys[monkey.Truecase]
					m.Levels = append(m.Levels, next)
					monkeys[monkey.Truecase] = m
				} else {
					m := monkeys[monkey.Falsecase]
					m.Levels = append(m.Levels, next)
					monkeys[monkey.Falsecase] = m
				}
			}
			monkey.Levels = []*big.Int{}
			monkeys[i] = monkey
		}
	}
	result := []Monkey{}
	for _, v := range monkeys {
		result = append(result, v)
	}
	sort.Sort(ByInspect(result))
	return result[len(result)-1].Inspect * result[len(result)-2].Inspect
}
