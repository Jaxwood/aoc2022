package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestDay14a(t *testing.T) {
	actual := day14(input)
	expected := 24
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}