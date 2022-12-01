package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Calorie struct {
	calories []float64
}

func (c Calorie) sum() float64 {
	var sum = 0.0

	for _, calorie := range c.calories {
		sum += calorie
	}

	return sum
}

func read(filename string) []Calorie {
	readFile, err := os.Open(filename)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []float64
	var calories []Calorie

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			calories = append(calories, Calorie{calories: fileLines})
			fileLines = make([]float64, 0)
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
		}
		fileLines = append(fileLines, num)
	}
	calories = append(calories, Calorie{calories: fileLines})

	return calories
}

func day01a(filename string) float64 {

	calories := read(filename)
	var best = 0.0
	for _, calorie := range calories {
		best = math.Max(best, calorie.sum())
	}

	return best
}

func day01b(filename string) float64 {

	calories := read(filename)
	var sumOfCalories []float64

	for _, calorie := range calories {
		sumOfCalories = append(sumOfCalories, calorie.sum())
	}

	sort.Float64s(sumOfCalories)

	sum := 0.0
	for _, num := range sumOfCalories[len(sumOfCalories)-3:] {
		sum += num
	}
	return sum
}
