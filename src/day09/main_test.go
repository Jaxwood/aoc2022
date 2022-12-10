package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestDay09a(t *testing.T) {
	actual := day09(input)
	expected := 13
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
