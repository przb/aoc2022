package day01

import (
	"fmt"
	"github.com/przb/aoc2022/util"
	"strconv"
)

const inputFilename = "input01-1.txt"
const inputDir = "inputs/"

type elf struct {
	foodCalories []int
	total        int
}

// / Returns an elf, from the input string. assumes that the entire string input is a single elf
func getElf(lines []string) (*elf, error) {
	var e elf
	e.total = 0
	for _, line := range lines {
		calorie, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		e.total += calorie
		e.foodCalories = append(e.foodCalories, calorie)
	}
	return &e, nil
}

func getElves(lines []string) ([]*elf, error) {
	elves := make([]*elf, 0)
	startI := 0
	for i, line := range lines {
		if len(line) == 0 { // if the line is empty
			elf, err := getElf(lines[startI:i])
			if err != nil {
				return nil, err
			}
			elves = append(elves, elf)
			startI = i + 1
		} else if i == len(lines)-1 { // the last line
			elf, err := getElf(lines[startI : i+1])
			if err != nil {
				return nil, err
			}
			elves = append(elves, elf)
			startI = i + 1
		}
	}
	return elves, nil
}

func HowFatIsTheFastestElf() {
	lines, err := util.ReadLines(inputDir + inputFilename)
	if err != nil {
		panic(err)
	}
	elves, err := getElves(lines)
	if err != nil {
		panic(err)
	}
	max := elves[0].total

	for _, elf := range elves {
		if elf.total > max {
			max = elf.total
		}
	}

	fmt.Println(max)

}
