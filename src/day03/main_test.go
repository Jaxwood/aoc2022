package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day03.txt
var filea string

func TestDay03a(t *testing.T) {
	actual := day03a(input)
	expected := 157
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay03b(t *testing.T) {
	actual := day03a(filea)
	expected := 8394
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay03c(t *testing.T) {
	actual := day03b(input)
	expected := 70
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
