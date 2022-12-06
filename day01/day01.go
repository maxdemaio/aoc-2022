package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(data string) int {
	groups := strings.Split(data, "\n\n")
	var elfCalCount = [3]int{0, 0, 0}

	for _, group := range groups {
		sum := 0
		for _, line := range strings.Split(group, "\n") {
			result, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += result
		}
		if sum > elfCalCount[0] {
			elfCalCount[2] = elfCalCount[1]
			elfCalCount[1] = elfCalCount[0]
			elfCalCount[0] = sum
		} else if sum > elfCalCount[1] {
			elfCalCount[2] = elfCalCount[1]
			elfCalCount[1] = sum
		} else if sum > elfCalCount[2] {
			elfCalCount[2] = sum
		}
	}

	return elfCalCount[0]
}

func Part2(data string) int {
	groups := strings.Split(data, "\n\n")
	var elfCalCount = [3]int{0, 0, 0}

	for _, group := range groups {
		sum := 0
		for _, line := range strings.Split(group, "\n") {
			result, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += result
		}
		if sum > elfCalCount[0] {
			elfCalCount[2] = elfCalCount[1]
			elfCalCount[1] = elfCalCount[0]
			elfCalCount[0] = sum
		} else if sum > elfCalCount[1] {
			elfCalCount[2] = elfCalCount[1]
			elfCalCount[1] = sum
		} else if sum > elfCalCount[2] {
			elfCalCount[2] = sum
		}
	}

	return elfCalCount[0] + elfCalCount[1] + elfCalCount[2]
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
