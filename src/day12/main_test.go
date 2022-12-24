package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day12.txt
var file string

func TestDay12a(t *testing.T) {
	actual := day12(input)
	expected := 31
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay12b(t *testing.T) {
	actual := day12(file)
	expected := 391
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
