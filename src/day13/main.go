package main

import (
	"encoding/json"
	"fmt"
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

// 0 if left == right
// -1 if left < right
// +1 if left > right
func compare(left interface{}, right interface{}) int {
	// compare each element
	leftSide, leftSideOk := left.([]interface{})
	rightSide, rightSideOk := right.([]interface{})
	if leftSideOk && rightSideOk {
		for i := 0; i < len(leftSide); i++ {
			// right ran out of items
			if len(rightSide) == i {
				return 1
			}
			if compare(leftSide[i], rightSide[i]) == 1 {
				return 1
			}
		}
	} else if leftSideOk {
		rightNum := right.(float64)
		fmt.Println(rightNum)
		if compare(left, []interface{}{rightNum}) == 1 {
			return 1
		}
	} else if rightSideOk {
		leftNum := left.(float64)
		fmt.Println(leftNum)
	} else {
		rightNum := right.(float64)
		leftNum := left.(float64)
		fmt.Println("comparing", leftNum, rightNum)
		if leftNum > rightNum {
			return +1
		}
		if leftNum < rightNum {
			return -1
		}
		return 0
	}
	return +1
}

func convert(str string) interface{} {
	var result interface{}
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		panic("could not parse json")
	}
	return result
}

func day13(file string) int {
	pairs := parse(file)
	sum := 0
	for idx, pair := range pairs {
		if compare(pair.Left, pair.Right) == +1 {
			sum += idx + 1
		}
	}
	return sum
}
