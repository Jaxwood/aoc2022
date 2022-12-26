package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestDay16a(t *testing.T) {
	actual := day16(input)
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
