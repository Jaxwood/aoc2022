package main

import (
	"strconv"
	"strings"
)

type Pair struct {
	Left  string
	Right string
}

func parse(file string) []Pair {
	result := []Pair{}

	lines := strings.Split(file, "\n")
	pairs := [][]string{}
	pair := []string{}
	for _, line := range lines {
		if line == "" {
			pairs = append(pairs, pair)
			pair = []string{}
		} else {
			pair = append(pair, line)
		}
	}

	pairs = append(pairs, pair)

	for _, p := range pairs {
		result = append(result, Pair{p[0], p[1]})
	}

	return result
}

func compare(left []interface{}, right []interface{}) bool {
	// compare each element
	for i := 0; i < len(left); i++ {
		// no more right elements
		if i == len(right) {
			return false
		}
		leftInt, leftIntOk := left[i].(int)
		rightInt, rightIntOk := right[i].(int)
		if leftIntOk && rightIntOk {
			// equal - check next pair
			if leftInt == rightInt {
				continue
			}
			// compare left and right
			return leftInt < rightInt
		} else if leftIntOk {
			if !compare([]interface{} {leftInt}, right[i].([]interface{})) {
				return false
			}
		} else if rightIntOk {
			if !compare(left[i].([]interface{}), []interface{} { rightInt }) {
				return false
			}
		} else {
			if !compare(left[i].([]interface{}), right[i].([]interface{})) {
				return false
			}
		}
	}
	return true
}

func convert(str string) []interface{} {
	queue := str
	var result [][]interface{}
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		switch next {
			case '[':
				result = append(result, []interface{}{})
			case ']':
				idx := len(result)-1
				tmp := result[idx]
				if idx > 0 {
					result = result[0:idx]
					result[idx-1] = append(result[idx-1], tmp)
				}
			case ',':
				continue
			default:
				num := string(next)
				for {
					if _, ok := strconv.Atoi(string(queue[0])); ok == nil {
						num += string(queue[0])
						queue = queue[1:]
					} else {
						break
					}
				}
				idx := len(result)-1
				val, _ := strconv.Atoi(string(num))
				result[idx] = append(result[idx], val)
		}
	}
	return result[0]
}

func day13(file string) int {
	pairs := parse(file)
	sum := 0
	for idx, pair := range pairs {
		if compare(convert(pair.Left), convert(pair.Right)) {
			sum += idx + 1
		}
	}
	return sum
}
