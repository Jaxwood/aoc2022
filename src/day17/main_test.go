package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day17.txt
var file string

func TestDay17a(t *testing.T) {
	actual := day17(input)
	expected := 3068
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
