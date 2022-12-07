package main

import "testing"

var testStr string = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	result := Part1(testStr)
	expected := 15
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testStr)
	expected := 12
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
