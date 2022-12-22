package day04

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "inputs/input04-1.txt"

func parsePair(pair string) (int, int) {
	stringsPair := strings.Split(pair, "-")
	num1, err := strconv.Atoi(stringsPair[0])
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(stringsPair[1])
	if err != nil {
		panic(err)
	}
	return num1, num2
}

func completelyOverlaps(e1Min, e1Max, e2Min, e2Max int) bool {
	if e1Min <= e2Min && e1Max >= e2Max {
		return true
	} else if e2Min <= e1Min && e2Max >= e1Max {
		return true
	}
	return false
}

func partiallyOverlaps(e1Min, e1Max, e2Min, e2Max int) bool {
	if e1Min <= e2Min && e1Max >= e2Min {
		return true
	} else if e2Min <= e1Min && e2Max >= e1Min {
		return true
	}
	return false
}

func NumOverlaps() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	numOverlaps := 0
	for _, line := range records {
		elf1Min, Elf1Max := parsePair(line[0])
		elf2Min, Elf2Max := parsePair(line[1])
		if partiallyOverlaps(elf1Min, Elf1Max, elf2Min, Elf2Max) {
			numOverlaps++
		}
	}
	fmt.Println(numOverlaps)
}
