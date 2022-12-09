package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type vector struct {
	pair
}

type point struct {
	pair
}

type pair struct {
	X int
	Y int
}

func move(direction string, amount int) {

}

func readfileVectors(filename string) []vector {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	for filescanner.Scan() {
		lines = append(lines, filescanner.Text())
	}

	vectors := make([]vector, 0)
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		value, _ := strconv.Atoi(instruction[1])

		if instruction[0] == "D" {
			vectors = append(vectors, vector{pair{X: 0, Y: -1 * value}})
		} else if instruction[0] == "U" {
			vectors = append(vectors, vector{pair{X: 0, Y: value}})

		} else if instruction[0] == "L" {
			vectors = append(vectors, vector{pair{X: -1 * value, Y: 0}})
		} else if instruction[0] == "R" {
			vectors = append(vectors, vector{pair{X: value, Y: 0}})
		}
	}
	return vectors
}

func main() {
	filename := "input.txt"
	readfileVectors(filename)

}
