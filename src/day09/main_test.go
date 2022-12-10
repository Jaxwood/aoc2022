package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day09.txt
var file string

func TestDay09a(t *testing.T) {
	actual := day09(input)
	expected := 13
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay09b(t *testing.T) {
	actual := day09(file)
	expected := 6087
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
