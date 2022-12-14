package day03

import (
	"fmt"
	"github.com/przb/aoc2022/util"
	"strings"
)

const filename = "inputs/input03-1.txt"

func getDuplicatePriority(line string) int {

	halfIndex := len(line) / 2
	firstHalf := strings.Split(line[:halfIndex], "")
	secondHalf := strings.Split(line[halfIndex:], "")

	// Technically bad bigO runtime
	for _, l := range firstHalf {
		if strings.Contains(line[halfIndex:], l) {
			return getPriority(l)
		}
	}

	for _, l := range secondHalf {
		if strings.Contains(line[halfIndex:], l) {
			return getPriority(l)
		}
	}
	return -1

}

// treats input s as a single character string
func getPriority(s string) int {
	a := 'a'
	A := 'A'
	if rune(s[0]) >= a { // 'A' is 65 and 'a' is 97
		return int(rune(s[0])) - int(a) + 1
	} else {
		return int(rune(s[0])) - int(A) + 26 + 1
	}

}

func findUnion(e1, e2 string) string {
	badgeCandidates := ""
	e1Arr := strings.Split(e1, "")
	for _, letter := range e1Arr {
		if strings.Contains(e2, letter) && !strings.Contains(badgeCandidates, letter) {
			badgeCandidates += letter
		}
	}
	return badgeCandidates
}

func GetBadges() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	total := 0

	for i := 2; i < len(lines); i += 3 {
		badgeCandidnate := findUnion(lines[i-2], lines[i-1])
		badgeCandidnate = findUnion(badgeCandidnate, lines[i])
		if len(badgeCandidnate) != 1 {
			fmt.Println("no single union found")
		} else {
			total += getPriority(badgeCandidnate)
		}
	}
	fmt.Println("total: ", total)
}

func GetItems() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	total := 0

	for _, line := range lines {
		total += getDuplicatePriority(line)
	}
	fmt.Println("total: ", total)
}
