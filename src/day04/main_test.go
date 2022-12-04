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
	actual := day04a(input)
	expected := 2
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay04b(t *testing.T) {
	actual := day04a(filea)
	expected := 567
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay04c(t *testing.T) {
	actual := day04b(input)
	expected := 4
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay04d(t *testing.T) {
	actual := day04b(filea)
	expected := 907
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
