package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	forest := [][]int{}

	for scanner.Scan() {
		templ := strings.Split(scanner.Text(), "")

		line := []int{}
		for _, v := range templ {
			value, _ := strconv.Atoi(v)
			line = append(line, value)
		}

		forest = append(forest, line)
	}

	log.Println(forest)

	n := len(forest)
	// calculate perimeter of a square 2 less than height (inner)
	// add back in 4 corners
	// avoids double counting corners
	sumv := 4*(n-2) + 4
	log.Println(sumv)

	for r := 1; r < n-1; r++ {
		for c := 1; c < n-1; c++ {
			visible := []int{1, 1, 1, 1}
			for i := 0; i < c; i++ {
				if forest[r][i] >= forest[r][c] {
					visible[0] = 0
					break
				}
			}

			for i := c + 1; i < n; i++ {
				if forest[r][i] >= forest[r][c] {
					visible[1] = 0
					break
				}
			}

			for i := 0; i < r; i++ {
				if forest[i][c] >= forest[r][c] {
					visible[2] = 0
				}
			}

			for i := r + 1; i < n; i++ {
				if forest[i][c] >= forest[r][c] {
					visible[3] = 0
				}
			}

			for _, v := range visible {
				if v == 1 {
					log.Println(forest[r][c], visible)
					sumv++
					break
				}
			}
		}
	}

	return sumv
}

func Part2(filepath string) int {
	/*
		Notes:
			01234

		0	xxxxx
		1	xx5xx
		2	xx5xx
		3	xx5xx
		4	xxxxx

		- for example let's take the bottom most 5
		- the max top scenic score is 3
		- when incrementing from top down we encounter a 5 at pos 1
		- scenic score is now 2 (3-1)
		- we encounter the next 5 at pos 2
		- scenic score is now only 1 (3-2)

		we rinse and repeat for all directions
	*/
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	forest := [][]int{}

	for scanner.Scan() {
		templ := strings.Split(scanner.Text(), "")

		line := []int{}
		for _, v := range templ {
			value, _ := strconv.Atoi(v)
			line = append(line, value)
		}

		forest = append(forest, line)
	}

	log.Println(forest)

	n := len(forest)
	maxScenicScore := 0
	log.Println(maxScenicScore)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			maxScenicScores := []int{r, n - c - 1, n - r - 1, c}
			scenicScores := []int{0, 0, 0, 0}

			// tree found to the top that's taller or equal
			// (approach: top down)
			// update top scenic score
			tblock := false
			for i := 0; i < r; i++ {
				if forest[i][c] >= forest[r][c] {
					scenicScores[0] = maxScenicScores[0] - i
					tblock = true
				}
			}
			if !tblock {
				scenicScores[0] = maxScenicScores[0]
			}

			// tree found to the right that's taller or equal
			// (approach: right one to edge)
			// update right scenic score
			rblock := false
			for i := c + 1; i < n; i++ {
				if forest[r][i] >= forest[r][c] {
					scenicScores[1] = i - n + maxScenicScores[1] + 1
					rblock = true
					break
				}
			}
			if !rblock {
				scenicScores[1] = maxScenicScores[1]
			}

			// tree found to the bottom that's taller or equal
			// (approach: below one to edge)
			// update bottom scenic score
			bblock := false
			for i := r + 1; i < n; i++ {
				if forest[i][c] >= forest[r][c] {
					scenicScores[2] = i - n + maxScenicScores[2] + 1
					bblock = true
					break
				}
			}
			if !bblock {
				scenicScores[2] = maxScenicScores[2]
			}

			// tree found to the left that's taller or equal
			// (approach: left edge to tree)
			// update left scenic score
			lblock := false
			for i := 0; i < c; i++ {
				if forest[r][i] >= forest[r][c] {
					scenicScores[3] = maxScenicScores[3] - i
					lblock = true
				}
			}
			if !lblock {
				scenicScores[3] = maxScenicScores[3]
			}

			// calculate total scenic score (top right bottom left)
			product := 1
			for _, v := range scenicScores {
				product *= v
			}

			fmt.Println()
			fmt.Printf("scenic score for tree [%d, %d]: %d", r, c, product)
			fmt.Println()

			// if greater than previous scores, change
			if product > maxScenicScore {
				maxScenicScore = product
			}
		}
	}

	return maxScenicScore
}

func main() {
	// change byte[] to string
	fmt.Printf("part2: %d", Part2("input.txt"))
}
