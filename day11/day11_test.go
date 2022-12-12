package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := Run(20)
	expected := 10605
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
