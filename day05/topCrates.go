package day05

import (
	"fmt"
	"github.com/przb/aoc2022/util"
	"strconv"
	"strings"
)

const (
	filename = "inputs/input05.txt"

	length = 9

	amountIndex = 1
	fromIndex   = 3
	toIndex     = 5
)

/*
   , [H],    ,    , [H],    ,    , [V],
   , [V],    ,    , [V], [J],    , [F], [F]
   , [S], [L],    , [M], [B],    , [L], [J]
   , [C], [N], [B], [W], [D],    , [D], [M]
[G], [L], [M], [S], [S], [C],    , [T], [V]
[P], [B], [B], [P], [Q], [S], [L], [H], [B]
[N], [J], [D], [V], [C], [Q], [Q], [M], [P]
[R], [T], [T], [R], [G], [W], [F], [W], [L]
 1 ,  2 ,  3 ,  4 ,  5 ,  6 ,  7 ,  8 ,  9
*/

// returns a 2-d array of slices. This represents an array of stacks, where the stacks[0] is the first stack, with the
// item on the top of the stack at index len(stacks[0]) - 1
func parseStart(start []string) [length][]string {

	stacks := [length][]string{}
	lines := make([][]string, 0)
	for i := len(start) - 1; i >= 0; i-- {
		line := start[i]
		newLine := strings.Split(line, ",")
		for i, item := range newLine {
			newLine[i] = strings.TrimSpace(item)
			if len(newLine[i]) != 0 {
				stacks[i] = append(stacks[i], newLine[i])
			}
		}
		lines = append(lines, newLine)
	}

	return stacks
}

func performMoves(stacks *[length][]string, moves []string) {
	for _, line := range moves {
		tokens := strings.Split(line, " ")
		amount, err := strconv.Atoi(tokens[amountIndex])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(tokens[fromIndex])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(tokens[toIndex])
		if err != nil {
			panic(err)
		}
		moveMultipleCrates(stacks, amount, from-1, to-1)
	}
}

// Moves amount items from the stack indexed at from to the stack indexed at to. Does this one item at a time
func moveCrates(stacks *[length][]string, amount, from, to int) {
	for i := 0; i < amount; i++ {
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
		stacks[from] = stacks[from][:len(stacks[from])-1]
	}
}

// Moves amount items from the stack indexed at from to the stack indexed at to. Retians the order of the moved items
func moveMultipleCrates(stacks *[length][]string, amount, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
	stacks[from] = stacks[from][:len(stacks[from])-amount]
}

func printTops(stacks [length][]string) {
	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1][1])
	}
	fmt.Printf("\n")
}

func TopCrates() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	stacks := parseStart(lines[:8])
	performMoves(&stacks, lines[10:])

	printTops(stacks)
}
