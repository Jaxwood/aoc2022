package main

import (
	_ "embed"
	"testing"
)

//go:embed day01.txt
var filea string

//go:embed day01a.txt
var fileb string

func TestDay01a(t *testing.T) {
	actual := day01a(filea)
	expected := 24000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01aa(t *testing.T) {
	actual := day01a(fileb)
	expected := 69310.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01b(t *testing.T) {
	actual := day01b(filea)
	expected := 45000.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay01bb(t *testing.T) {
	actual := day01b(fileb)
	expected := 206104.0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
