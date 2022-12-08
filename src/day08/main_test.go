package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day08.txt
var file string

func TestDay08a(t *testing.T) {
	actual := day08(input)
	expected := 21
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay08b(t *testing.T) {
	actual := day08(file)
	expected := 1684
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay08c(t *testing.T) {
	actual := day08b(input)
	expected := 8.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay08d(t *testing.T) {
	actual := day08b(file)
	expected := 486540.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
