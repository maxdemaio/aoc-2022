package main

import (
	"fmt"
	"log"
	"os"
)

func allDiff(s string) bool {
	// allocate array of size of set of ascii chars
	var tab ['z' + 1]bool
	for i := 0; i < len(s); i++ {
		// duplicate found, not unique
		if tab[s[i]] {
			return false
		}
		// flip boolean in set
		tab[s[i]] = true
	}
	return true
}

func Part1(data string, n int) int {
	for i := 0; i < len(data)-n; i++ {
		s := data[i : i+n]
		if allDiff(s) {
			// found a unique set of n
			return i + n
		}
	}

	// no unique set of n found, return 1
	return 1
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// change byte[] to string
	data_s := string(data)
	fmt.Println("part1: ", Part1(data_s, 4))
	fmt.Println("part2: ", Part1(data_s, 14))
}
