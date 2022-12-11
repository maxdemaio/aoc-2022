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

type Pos struct {
	x, y int
}

/*
State to contain the following:
- Number of knots
- Positions of knots (Oth position is for the head and so on)
- Set of positions where the tail knots have been
*/
type State struct {
	n    int
	rope []Pos
	path utils.Set[Pos]
}

// Return a pointer to a created State type in memory
func newState(n int) *State {
	res := &State{
		n:    n,
		rope: make([]Pos, n),
		path: utils.MakeSet[Pos](),
	}
	// Add starting position as default in set
	res.path.Add(res.rope[n-1])
	return res
}

// Method with receiver argument to act upon pointer to State type
// This is like class methods in OOP languages. Go does not have classes.
func (s *State) move(dir string) {
	// Move head of rope based on current direction
	switch dir {
	case "U":
		s.rope[0].y++
	case "D":
		s.rope[0].y--
	case "L":
		s.rope[0].x--
	case "R":
		s.rope[0].x++
	}
}

// Method with receiver argument to act upon pointer to State type
// This is like class methods in OOP languages. Go does not have classes.
func (s *State) moveTail() {
	// Each knot in the rope will follow the one in front of it (i-1)
	// If either x/y absolute difference between those knots is > 1
	// Adjust the knot to be within 1 distance to the knot in front of itself (i-1)
	for i := 1; i < s.n; i++ {
		delta := Pos{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
		if utils.Abs(delta.x) <= 1 && utils.Abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			s.rope[i].y++
		} else if delta.y < 0 {
			s.rope[i].y--
		}
		if delta.x > 0 {
			s.rope[i].x++
		} else if delta.x < 0 {
			s.rope[i].x--
		}
	}
	s.path.Add(s.rope[s.n-1])
}

func Run(filepath string, n int) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := newState(n)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		amt, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Println(err)
			return -2
		}

		for i := 0; i < amt; i++ {
			state.move(dir)
			state.moveTail()
		}
	}

	return state.path.Len()
}

func main() {
	fmt.Printf("part1: %d", Run("input.txt", 2))
	fmt.Println()
	fmt.Printf("part2: %d", Run("input.txt", 10))
}
