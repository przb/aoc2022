package main

import (
	"bufio"
	"fmt"
	"github.com/przb/aoc2022/util"
	"os"
	"strconv"
	"strings"
)

func move(visitedPoints map[util.Point]int, direction util.Vector, tail, head *util.Point) (util.Point, util.Point) {
	head = head.Add(direction)
	difX := head.X - tail.X
	difY := head.Y - tail.Y
	for !tail.Adjacent(head) && !tail.Equals(head) {
		var unitVector util.Vector
		if !tail.SameRow(head) && !tail.SameColumn(head) {
			X := difX / util.IntAbs(difX)
			Y := difY / util.IntAbs(difY)
			unitVector = util.Vector{X, Y}
			// then move diagonally
		} else if tail.SameColumn(head) {
			// move the y direction
			Y := difY / util.IntAbs(difY)
			unitVector = util.Vector{0, Y}
		} else if tail.SameRow(head) {
			// move the X direction
			X := difX / util.IntAbs(difX)
			unitVector = util.Vector{X, 0}

		}
		tail = tail.Add(unitVector)
		visitedPoints[util.Point{tail.X, tail.Y}]++
	}
	return *head, *tail
}

func readfileVectors(filename string) []util.Vector {
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

	vectors := make([]util.Vector, 0)
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		value, _ := strconv.Atoi(instruction[1])

		if instruction[0] == "D" {
			vectors = append(vectors, util.Vector{X: 0, Y: -1 * value})
		} else if instruction[0] == "U" {
			vectors = append(vectors, util.Vector{X: 0, Y: value})
		} else if instruction[0] == "L" {
			vectors = append(vectors, util.Vector{X: -1 * value, Y: 0})
		} else if instruction[0] == "R" {
			vectors = append(vectors, util.Vector{X: value, Y: 0})
		}
	}
	return vectors
}

func main() {
	filename := "input.txt"
	visitedPoints := map[util.Point]int{util.Point{0, 0}: 1}
	tail := util.Point{0, 0}
	head := util.Point{0, 0}
	vectors := readfileVectors(filename)
	for _, v := range vectors {
		head, tail = move(visitedPoints, v, &tail, &head)
	}
	fmt.Println(len(visitedPoints))
}
