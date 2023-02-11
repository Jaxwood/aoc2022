package main

import (
	_ "embed"
	"testing"
)

type TestCase struct {
	Left     []int
	Right    []int
	Expected bool
}

//go:embed input.txt
var input string

//go:embed day13.txt
var file string

func TestDay13a(t *testing.T) {
	actual := day13(input)
	expected := 13
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func _TestDay13b(t *testing.T) {
	actual := day13(file)
	expected := 4437 // too low
	// 5289 too high
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareA(t *testing.T) {
	actual := compare(convert("[1,1,40,1,1]"), convert("[1,1,50,1,1]"))
	expected := 1
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareB(t *testing.T) {
	actual := compare(convert("[[1],[2,3,4]]"), convert("[[1],4]"))
	expected := 1
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareC(t *testing.T) {
	actual := compare(convert("[9]"), convert("[[8,7,6]]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareD(t *testing.T) {
	actual := compare(convert("[[4,4],4,4]"), convert("[[4,4],4,4,4]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareE(t *testing.T) {
	actual := compare(convert("[7,7,7,7]"), convert("[7,7,7]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareF(t *testing.T) {
	actual := compare(convert("[[[]]]"), convert("[[]]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareG(t *testing.T) {
	actual := compare(convert("[1,[2,[3,[4,[5,6,7]]]],8,9]"), convert("[1,[2,[3,[4,[5,6,0]]]],8,9]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}

func TestCompareI(t *testing.T) {
	actual := compare(convert("[[[4,4],4,4]]"), convert("[[4,5],2,3,4]"))
	expected := 0
	if actual != expected {
		t.Fatalf(`actual = %v, expected = %v`, actual, expected)
	}
}
