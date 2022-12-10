package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day10.txt
var file string

func TestDay10a(t *testing.T) {
	actual := day10(input)
	expected := 13140
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay10b(t *testing.T) {
	actual := day10(file)
	expected := 12460
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay10c(t *testing.T) {
	actual := day10b(input)
	expected := "##..##..##..##..##..##..##..##..##..##..###...###...###...###...###...###...###.####....####....####....####....####....#####.....#####.....#####.....#####.....######......######......######......###########.......#######.......#######....."
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay10d(t *testing.T) {
	actual := day10b(file)
	expected := "####.####.####.###..###...##..#..#.#....#.......#.#....#..#.#..#.#..#.#.#..#....###....#..###..#..#.#..#.#..#.##...#....#.....#...#....###..###..####.#.#..#....#....#....#....#....#.#..#..#.#.#..#....####.####.#....#....#..#.#..#.#..#.####."
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
