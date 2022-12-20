package day02

import (
	"fmt"
	"github.com/przb/aoc2022/util"
)

const (
	mustLose = "X"
	mustTie  = "Y"
	mustWin  = "Z"
)

func linePts(line []string) int {
	// Rock
	if line[0] == opponentRock && line[1] == mustLose {
		return scissorsPts + losePts
	}
	if line[0] == opponentRock && line[1] == mustTie {
		return rockPts + tiePts
	}
	if line[0] == opponentRock && line[1] == mustWin {
		return paperPts + winPts
	}

	// Paper
	if line[0] == opponentPaper && line[1] == mustLose {
		return rockPts + losePts
	}
	if line[0] == opponentPaper && line[1] == mustTie {
		return paperPts + tiePts
	}
	if line[0] == opponentPaper && line[1] == mustWin {
		return scissorsPts + winPts
	}

	// Scissors
	if line[0] == opponentScissors && line[1] == mustLose {
		return paperPts + losePts
	}
	if line[0] == opponentScissors && line[1] == mustTie {
		return scissorsPts + tiePts
	}
	if line[0] == opponentScissors && line[1] == mustWin {
		return rockPts + winPts
	}
	return 0
}

func TotalPtsPt2() {
	lines, err := util.ReadTokens(filename)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range lines {
		sum += linePts(line)
	}
	fmt.Println(sum)
}
