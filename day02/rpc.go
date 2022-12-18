package day02

import (
	"fmt"
	"github.com/przb/aoc2022/util"
)

const (
	filename = "inputs/input02-1.txt"

	losePts = 0
	tiePts  = 3
	winPts  = 6

	rockPts     = 1
	paperPts    = 2
	scissorsPts = 3

	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
	myRock           = "X"
	myPaper          = "Y"
	myScissors       = "Z"
)

// Returns the sum of the points you get from choosing rock, paper, or scissors
func selectionSum() int {
	sum := 0
	m := map[string]int{myRock: rockPts, myPaper: paperPts, myScissors: scissorsPts}

	lines, err := util.ReadTokens(filename)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		sum += m[line[1]]
	}

	return sum
}

func win(opponentSel, mySel string) bool {
	return (opponentSel == opponentRock && mySel == myPaper) ||
		(opponentSel == opponentPaper && mySel == myScissors) ||
		(opponentSel == opponentScissors && mySel == myRock)
}

// Returns the points you get from the rounds if you win
func winSum() int {

	// if equal -> 3 pts
	// if won -> 6 pts
	// if lost -> 0 pts
	sum := 0

	lines, err := util.ReadTokens(filename)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		if line[0] == line[1] {
			sum += tiePts
		} else if win(line[0], line[1]) {
			sum += winPts
		} else {
			sum += losePts
		}
	}

	return sum
}

func TotalPts() {
	fmt.Println(selectionSum() + winSum())
}
