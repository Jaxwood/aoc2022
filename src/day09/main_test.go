package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed input2.txt
var input2 string

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

func TestDay09c(t *testing.T) {
	actual := day09b(input)
	expected := 1
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay09d(t *testing.T) {
	actual := day09b(input2)
	expected := 36
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay09e(t *testing.T) {
	actual := day09b(file)
	expected := 2493
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
