package main

import (
	_ "embed"
	"testing"
)

func TestDay01a(t *testing.T) {
	actual := day01a("./day01.txt")
	expected := 24000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01aa(t *testing.T) {
	actual := day01a("./day01a.txt")
	expected := 69310.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01b(t *testing.T) {
	actual := day01b("./day01.txt")
	expected := 45000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01bb(t *testing.T) {
	actual := day01b("./day01a.txt")
	expected := 206104.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
