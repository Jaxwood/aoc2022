package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day05.txt
var filea string

func TestDay05a(t *testing.T) {
	initial := map[int][]rune{
		1: []rune{'Z', 'N'},
		2: []rune{'M', 'C', 'D'},
		3: []rune{'P'},
	}
	actual := day05a(input, Cargo{initial})
	expected := "CMZ"
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay05b(t *testing.T) {
	initial := map[int][]rune{
		1: []rune{'H', 'C', 'R'},
		2: []rune{'B', 'J', 'H', 'L', 'S', 'F'},
		3: []rune{'R', 'M', 'D', 'H', 'J', 'T', 'Q'},
		4: []rune{'S', 'G', 'R', 'H', 'Z', 'B', 'J'},
		5: []rune{'R', 'P', 'F', 'Z', 'T', 'D', 'C', 'B'},
		6: []rune{'T', 'H', 'C', 'G'},
		7: []rune{'S', 'N', 'V', 'Z', 'B', 'P', 'W', 'L'},
		8: []rune{'R', 'J', 'Q', 'G', 'C'},
		9: []rune{'L', 'D', 'T', 'R', 'H', 'P', 'F', 'S'},
	}
	actual := day05a(filea, Cargo{initial})
	expected := "SHQWSRBDL"
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
