package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed day11.txt
var file string

func TestDay11a(t *testing.T) {
	actual := day11a(input, 20)
	expected := int64(10605)
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay11b(t *testing.T) {
	actual := day11a(file, 20)
	expected := int64(182293)
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay11c(t *testing.T) {
	actual := day11a(input, 10000)
	expected := int64(2713310158)
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
