package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day07.txt
var file string

func TestDay07a(t *testing.T) {
	actual := day07(input)
	expected := 95437
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay07b(t *testing.T) {
	actual := day07(file)
	expected := 1141028
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay07c(t *testing.T) {
	actual := day07b(input)
	expected := 24933642
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay07d(t *testing.T) {
	actual := day07b(file)
	expected := 8278005
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
