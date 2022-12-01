package main

import (
	_ "embed"
	"testing"
)

func TestDay01(t *testing.T) {
    actual := day01a("./day01.txt")
    expected := 24000.0
    if actual != expected {
        t.Fatalf(`actual = %v, expected = %v`, actual, expected)
    }
}

func TestDay01a(t *testing.T) {
    actual := day01a("./day01a.txt")
    expected := 69310.0
    if actual != expected {
        t.Fatalf(`actual = %v, expected = %v`, actual, expected)
    }
}
