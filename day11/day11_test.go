package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := Run(1, 20)
	expected := 10605
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	// m.worryLevel = m.worryLevel % (23 * 19 * 13 * 17)
	result := Run(2, 10000)
	expected := 2713310158
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
