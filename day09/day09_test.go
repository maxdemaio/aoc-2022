package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	result := Part1("test.txt", 2, mvMap)
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestMain(m *testing.M) {
	// initialize the map in this function
	mvMap = make(map[string][2]int, 4)
	mvMap["R"] = [2]int{1, 0}
	mvMap["U"] = [2]int{0, 1}
	mvMap["L"] = [2]int{-1, 0}
	mvMap["D"] = [2]int{0, -1}

	// run the tests
	code := m.Run()

	// do any cleanup here, if necessary

	// exit with the same code as the tests
	os.Exit(code)
}
