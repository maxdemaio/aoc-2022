package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// stack of folders
	fs := []string{}

	// folder to folder filesize mappings
	m := map[string]int{}

	// for duplicate names
	// we use counter to make the folder names unique
	counter := 0

	for scanner.Scan() {
		// grab tokens
		fields := strings.Fields(scanner.Text())
		if fields[0] == "$" {
			if fields[1] == "cd" {
				if fields[2] != ".." {
					dir := fields[2]
					// if folder already found make unique w/ counter
					if _, ok := m[fields[2]]; ok {
						dir += strconv.Itoa(counter)
						counter++
					}
					// add dir to folder stack
					fs = append(fs, dir)

				} else {
					// go back one folder, pop last dir from stack
					fs = fs[:len(fs)-1]
				}
			}
		} else {
			// we're on a file, add its file size to whole folder stack
			if fields[0] != "dir" {
				for i := 0; i < len(fs); i++ {
					size, _ := strconv.Atoi(fields[0])
					m[fs[i]] += size
				}
			}
		}

	}

	// current units of unused disk
	unused := 70000000 - m["/"]

	// amount of units to remove
	// to get us to 30mil free units
	smallest := 30000000 - unused

	// start min off as the biggest folder, "/"
	min := m["/"]

	for _, v := range m {
		if v >= smallest && v < min {
			min = v
		}
	}

	return min
}

func Part1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// stack of folders
	fs := []string{}

	// folder to folder filesize mappings
	m := map[string]int{}

	// for duplicate names
	// we use counter to make the folder names unique
	counter := 0

	for scanner.Scan() {
		// grab tokens
		fields := strings.Fields(scanner.Text())
		if fields[0] == "$" {
			if fields[1] == "cd" {
				if fields[2] != ".." {
					dir := fields[2]
					// if folder already found make unique w/ counter
					if _, ok := m[fields[2]]; ok {
						dir += strconv.Itoa(counter)
						counter++
					}
					// add dir to folder stack
					fs = append(fs, dir)

				} else {
					// go back one folder, pop last dir from stack
					fs = fs[:len(fs)-1]
				}
			}
		} else {
			// we're on a file, add its file size to whole folder stack
			if fields[0] != "dir" {
				for i := 0; i < len(fs); i++ {
					size, _ := strconv.Atoi(fields[0])
					m[fs[i]] += size
				}
			}
		}

	}

	sum := 0
	for _, v := range m {
		if v <= 100000 {
			sum += v
		}
	}

	return sum
}

func main() {
	// change byte[] to string
	fmt.Println("part1: ", Part1("input.txt"))
	fmt.Println("part2: ", Part2("input.txt"))
}
