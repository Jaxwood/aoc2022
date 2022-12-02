package main

import (
	_ "embed"
	"testing"
)

//go:embed day02.txt
var filea string

//go:embed day02a.txt
var fileb string

func TestDay02a(t *testing.T) {
	actual := day02a(filea)
	expected := 10404
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay02b(t *testing.T) {
	actual := day02a(fileb)
	expected := 15
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay02c(t *testing.T) {
	actual := day02b(filea)
	expected := 10334
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestDay02d(t *testing.T) {
	actual := day02b(fileb)
	expected := 12
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
