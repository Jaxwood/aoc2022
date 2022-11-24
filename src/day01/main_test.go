package main

import (
	_ "embed"
	"testing"
)

func TestAdd(t *testing.T) {
    actual := add(1,2)
    expected := 4
    if actual != expected {
        t.Fatalf(`actual = %v, expected = %v`, actual, expected)
    }
}
