package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt", 2)
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
