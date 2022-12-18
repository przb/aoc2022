package day02

import (
	"github.com/przb/aoc2022/util"
)

const filename = "inputs/input02-1.txt"

// Returns the sum of the points you get from choosing rock, paper, or scissors
func selectionSum() int {
	sum := 0
	m := map[string]int{"X": 1, "Y": 2, "Z": 3}

	lines, err := util.ReadTokens(filename)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		sum += m[line[1]]
	}

	return sum
}
