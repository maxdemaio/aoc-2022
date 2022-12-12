package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Monkey struct {
	items       []int
	worryLevel  int
	testDivisBy int
	trueMonky   int
	falseMonky  int
}

type MonkeyJamboree struct {
	monkys []Monkey
}

// Return a pointer to a created Monkey type in memory
func newMonkey(items []int, testDivis int, tMonky int, fMonky int) *Monkey {
	monky := &Monkey{
		items:       items,
		testDivisBy: testDivis,
		trueMonky:   tMonky,
		falseMonky:  fMonky,
	}
	return monky
}

// Return a pointer to a created MonkeyJamboree type in memory
func newMonkeyJamboree(monkys []Monkey) *MonkeyJamboree {
	monkyJam := &MonkeyJamboree{
		monkys: monkys,
	}
	return monkyJam
}

func (m *Monkey) updateWorryLevel(worryLevel int, op string, amt int) {
	// Move head of rope based on current direction
	switch op {
	case "+":
		m.worryLevel = m.worryLevel + amt
	case "*":
		m.worryLevel = m.worryLevel * amt
	}
	// TODO: Monkey gets bored, round down

}

func (m *Monkey) divisBy() bool {
	// TODO: test if m.worryLevel mod has no remainder
	return false
}

func (mj *MonkeyJamboree) throw(from int, to int) {

}

func Run(filepath string, n int) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("can not open file")
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var monkyJam MonkeyJamboree

	for scanner.Scan() {

		if err != nil {
			fmt.Println(err)
			return -2
		}

		// TODO populate MonkeyJamboree with new Monkeys
		fmt.Println(monkyJam)
	}

	// Play out n rounds
	for i := 0; i < n; i++ {
		fmt.Println()
		fmt.Printf("playing round %d", i)
		fmt.Println()
	}

	return 0
}

func main() {
	fmt.Printf("part1: %d", Run("input.txt", 20))
	fmt.Println()
	// fmt.Printf("part2: %d", Run("input.txt", 10))
}
