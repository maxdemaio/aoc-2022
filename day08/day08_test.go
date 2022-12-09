package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	expected := 21
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	result := Part2("test.txt")
// 	expected := 8
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }
