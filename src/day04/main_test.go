package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day04.txt
var filea string

func TestDay04a(t *testing.T) {
	actual := day04(input)
	expected := 2
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay04b(t *testing.T) {
	actual := day04(filea)
	expected := 567
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
