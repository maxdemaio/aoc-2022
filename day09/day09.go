package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/maxdemaio/aoc-2022/utils"
)

var mvMap map[string][2]int

func touching(x1 int, x2 int, y1 int, y2 int) bool {
	xdif := utils.AbsDiffInt(x1, x2)
	ydif := utils.AbsDiffInt(y1, y2)
	return xdif <= 1 && ydif <= 1
}

func Part1(filepath string, n int, mvMap map[string][2]int) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into a string and an integer
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		amt, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Println(err)
			return -2
		}

		fmt.Println()
		fmt.Printf("direction: %s | amount: %d", dir, amt)
		fmt.Println()
	}

	fmt.Println(mvMap["R"])
	return 0
}

func Part2(filepath string) int {
	return 0
}

func main() {
	// init move map of size 4 and add moves
	mvMap = make(map[string][2]int, 4)
	mvMap["R"] = [2]int{1, 0}
	mvMap["U"] = [2]int{0, 1}
	mvMap["L"] = [2]int{-1, 0}
	mvMap["D"] = [2]int{0, -1}

	fmt.Printf("part2: %d", Part1("input.txt", 2, mvMap))
}
