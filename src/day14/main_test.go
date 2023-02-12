package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day14.txt
var file string

func TestDay14a(t *testing.T) {
	actual := day14(input)
	expected := 24
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay14b(t *testing.T) {
	actual := day14(file)
	expected := 799
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay14c(t *testing.T) {
	actual := day14b(input)
	expected := 93
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay14d(t *testing.T) {
	actual := day14b(file)
	expected := 29076
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
