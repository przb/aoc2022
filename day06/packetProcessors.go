package day06

import (
	"fmt"
	"github.com/przb/aoc2022/util"
	"strings"
)

const filename = "inputs/input06.txt"

func uniqueInString(s []string, sub string) bool {
	found := false
	for _, l := range s {
		if l == sub && found == true {
			return false
		} else if l == sub {
			found = true
		}
	}
	return true
}
func noDuplicates(s []string) bool {
	for _, l := range s {
		if !uniqueInString(s, l) {
			return false
		}
	}
	return true
}

func findIndexOfFirstUniqueString(s []string, length int) int {
	for i := length; i < len(s); i++ {
		substr := s[i-length : i]
		if noDuplicates(substr) {
			return i
		}
	}
	return -1
}

func NumCharacters() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	characters := strings.Split(lines[0], "") // there is only one line for this puzzle
	fmt.Println(findIndexOfFirstUniqueString(characters, 14))
}
