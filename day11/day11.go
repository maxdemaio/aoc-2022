package main

import (
	"fmt"
)

type Monkey struct {
	items        []int
	worryLevel   int
	testDivisBy  int
	trueMonky    int
	falseMonky   int
	op           string
	opAmt        int
	inspectCount int
}

type MonkeyJamboree struct {
	monkys []*Monkey
}

// Return a pointer to a created MonkeyJamboree type in memory
func newMonkeyJamboree(monkys []*Monkey) *MonkeyJamboree {
	monkyJam := &MonkeyJamboree{
		monkys: monkys,
	}
	return monkyJam
}

func (m *Monkey) updateWorryLevel(part int) bool {
	// If a monkey is holding 0 items at the start of its turn, its turn ends
	// We will return true if update success, false if turn should end
	if len(m.items) == 0 {
		return false
	}

	// Monkey has started inspecting
	m.inspectCount++

	// Init worrry level
	m.worryLevel = m.items[0]

	// Update worry level based on operation
	switch m.op {
	case "+":
		m.worryLevel = m.worryLevel + m.opAmt
	case "*":
		m.worryLevel = m.worryLevel * m.opAmt
	case "sq":
		m.worryLevel = m.worryLevel * m.worryLevel
	}

	// Monkey gets bored, round down (default functionality of int division)
	if part == 1 {
		m.worryLevel = m.worryLevel / 3
	} else {
		m.worryLevel = m.worryLevel % (3 * 13 * 19 * 17 * 5 * 7 * 11 * 2)
	}
	return true
}

func (m *Monkey) divisBy() bool {
	// Test if worry level is divisible by Monkey's amount
	result := m.worryLevel % m.testDivisBy

	if result == 0 {
		// x is divisible by y
		return true
	} else {
		// x is not divisible by y
		return false
	}
}

func (mj *MonkeyJamboree) throw(from int, to int) {
	// We want to pop front of queue from current monkey
	// Then we will toss the worry level to another Monkey
	fromMonky := mj.monkys[from]
	toMonky := mj.monkys[to]

	// Throw item with worry level X from fromMonkey to toMonkey
	toMonky.items = append(toMonky.items, fromMonky.worryLevel)

	// Pop item from front of queue from fromMonkey
	fromMonky.items = fromMonky.items[1:]
}

func Run(part int, n int) int {
	// Test
	// monkys := []*Monkey{
	// 	{[]int{79, 98}, 0, 23, 2, 3, "*", 19, 0},
	// 	{[]int{54, 65, 75, 74}, 0, 19, 2, 0, "+", 6, 0},
	// 	{[]int{79, 60, 97}, 0, 13, 1, 3, "sq", 0, 0},
	// 	{[]int{74}, 0, 17, 0, 1, "+", 3, 0},
	// }

	// Input
	monkys := []*Monkey{
		{[]int{54, 98, 50, 94, 69, 62, 53, 85}, 0, 3, 2, 1, "*", 13, 0},
		{[]int{71, 55, 82}, 0, 13, 7, 2, "+", 2, 0},
		{[]int{77, 73, 86, 72, 87}, 0, 19, 4, 7, "+", 8, 0},
		{[]int{97, 91}, 0, 17, 6, 5, "+", 1, 0},
		{[]int{78, 97, 51, 85, 66, 63, 62}, 0, 5, 6, 3, "*", 17, 0},
		{[]int{88}, 0, 7, 1, 0, "+", 3, 0},
		{[]int{87, 57, 63, 86, 87, 53}, 0, 11, 5, 0, "sq", 0, 0},
		{[]int{73, 59, 82, 65}, 0, 2, 4, 3, "+", 6, 0},
	}

	monkyJam := newMonkeyJamboree(monkys)

	// Play out n rounds
	for i := 0; i < n; i++ {
		fmt.Println()
		fmt.Printf("playing round %d", i)
		fmt.Println()
		// Each Monkey needs to throw all of its items before
		// The next Monkey
		for j := 0; j < len(monkyJam.monkys); j++ {
			// If updateWorryLevel is false, skip to next Monkey
			// This is because that Monkey had no items
			currMonky := monkyJam.monkys[j]
			for currMonky.updateWorryLevel(part) {
				if currMonky.divisBy() {
					monkyJam.throw(j, currMonky.trueMonky)
				} else {
					monkyJam.throw(j, currMonky.falseMonky)
				}
			}
		}
	}

	twoMonkeyValues := []int{0, 0}

	for k := 0; k < len(monkyJam.monkys); k++ {
		monkyInspectCount := monkyJam.monkys[k].inspectCount
		if monkyInspectCount > twoMonkeyValues[0] {
			twoMonkeyValues[1] = twoMonkeyValues[0]
			twoMonkeyValues[0] = monkyInspectCount
		} else if monkyInspectCount > twoMonkeyValues[1] {
			twoMonkeyValues[1] = monkyInspectCount
		}
		fmt.Println(monkyJam.monkys[k].inspectCount)
	}

	return twoMonkeyValues[0] * twoMonkeyValues[1]
}

func main() {
	fmt.Printf("part1: %d", Run(1, 20))
	fmt.Println()
	fmt.Printf("part2: %d", Run(2, 10000))
}
