package main

import "testing"

var testStr string = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestPart1(t *testing.T) {
	result := Part1(testStr, 4)
	expected := 7
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part1(testStr, 14)
	expected := 19
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
