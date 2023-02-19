package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day15.txt
var file string

func TestDay15a(t *testing.T) {
	actual := day15(input, 10)
	expected := 26
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay15b(t *testing.T) {
	actual := day15(file, 2000000)
	expected := 4717631
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay15c(t *testing.T) {
	actual := day15b(input, 0, 20)
	expected := 56000011
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func _TestDay15d(t *testing.T) {
	actual := day15b(file, 0, 4000000)
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}