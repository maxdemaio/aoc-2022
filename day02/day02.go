package main

import (
	"fmt"
	"log"
	"os"
)

func Part1(data string) int {
	return 0
}

func Part2(data string) int {
	return 0
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// change byte[] to string
	data_s := string(data)
	fmt.Println("part1: ", Part1(data_s))
	fmt.Println("part2: ", Part2(data_s))
}
