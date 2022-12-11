package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := Run("test.txt", 2)
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Run("test2.txt", 10)
	expected := 36
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
