package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	X float64
	Y float64
}

type Reading struct {
	Beacon Coord
	Sensor Coord
}

func parse(file string) []Reading {
	readings := []Reading{}
	r := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range strings.Split(file, "\n") {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sensorX, _ := strconv.ParseFloat(match[1], 64)
			sensorY, _ := strconv.ParseFloat(match[2], 64)
 			beaconX, _ := strconv.ParseFloat(match[3], 64)
			beaconY, _ := strconv.ParseFloat(match[4], 64)
			readings = append(readings, Reading{
				Coord{beaconX, beaconY},
				Coord{sensorX, sensorY},
			})
		}
	}
	return readings
}

// find the coords that are within a certain manhattan distance
func (r *Reading) on(manhattanDistance float64, line float64) []Coord {
	var coords []Coord
	for x := r.Sensor.X - manhattanDistance; x <= r.Sensor.X+manhattanDistance; x++ {
		if math.Abs(r.Sensor.X-x)+math.Abs(r.Sensor.Y-line) <= manhattanDistance {
				coords = append(coords, Coord{X: x, Y: line})
		}
	}
	return coords
}

// find the coords that are within a certain manhattan distance
func (r *Reading) within(manhattanDistance float64) []Coord {
	var coords []Coord
	for y := r.Sensor.Y - manhattanDistance; y <= r.Sensor.Y+manhattanDistance; y++ {
		for x := r.Sensor.X - manhattanDistance; x <= r.Sensor.X+manhattanDistance; x++ {
			if math.Abs(r.Sensor.X-x)+math.Abs(r.Sensor.Y-y) <= manhattanDistance {
					coords = append(coords, Coord{X: x, Y: y})
			}
		}
	}
	return coords
}

func (r *Reading) manhattan() float64 {
	return math.Abs(r.Beacon.X-r.Sensor.X) + math.Abs(r.Beacon.Y-r.Sensor.Y)
}

func day15(file string, line float64) int {
	readings := parse(file)
	result := map[Coord]bool{}
	for _, reading := range readings {
		coords := reading.on(reading.manhattan(), line)
		for _, coord := range coords {
			result[coord] = true
		}
	}
	for _, reading := range readings {
		if result[reading.Beacon] {
			delete(result, reading.Beacon)
		}
	}
	return len(result)
}

func (r *Reading) filter(low float64, high float64) bool {
	if r.Beacon.X < low || r.Beacon.X > high {
		return false
	}

	if r.Beacon.Y < low || r.Beacon.Y > high {
		return false
	}

	return true
}

func (c *Coord) signal() float64 {
	return c.X * 4000000 + c.Y
}

func day15b(file string, low float64, high float64) int {
	readings := parse(file)
	result := map[Coord]bool{}
	for _, reading := range readings {
		if reading.filter(low, high) {
			coords := reading.within(reading.manhattan())
			for _, coord := range coords {
				result[coord] = true
			}
		}
	}

	// draw
	for y := -10.0; y < 25.0; y++ {
		line := ""
		for x := -25.0; x < 50.0; x++ {
			if x == 14 && y == 11 {
				line += "x"
			} else if result[Coord{x,y}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

	return 0
}
