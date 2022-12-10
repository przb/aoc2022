package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vector struct {
	X int
	Y int
}

type point struct {
	X int
	Y int
}

func (p *point) sameRow(other *point) bool {
	return p.Y == other.Y
}
func (p *point) sameColumn(other *point) bool {
	return p.X == other.X
}

func (p *point) adjacent(other *point) bool {
	return (IntAbs(p.X-other.X) == 1 && IntAbs(p.Y-other.Y) == 1) ||
		(p.sameColumn(other) && IntAbs(p.Y-other.Y) == 1) ||
		(p.sameRow(other) && IntAbs(p.X-other.X) == 1)
}

func (p *point) add(v vector) *point {

	return &point{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

func (p *point) equals(other *point) bool {
	return p.X == other.X && p.Y == other.Y
}

func IntAbs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func move(visitedPoints map[point]int, direction vector, tail, head *point) (point, point) {
	head = head.add(direction)
	difX := head.X - tail.X
	difY := head.Y - tail.Y
	for !tail.adjacent(head) && !tail.equals(head) {
		var unitVector vector
		if !tail.sameRow(head) && !tail.sameColumn(head) {
			X := difX / IntAbs(difX)
			Y := difY / IntAbs(difY)
			unitVector = vector{X, Y}
			// then move diagonally
		} else if tail.sameColumn(head) {
			// move the y direction
			Y := difY / IntAbs(difY)
			unitVector = vector{0, Y}
		} else if tail.sameRow(head) {
			// move the X direction
			X := difX / IntAbs(difX)
			unitVector = vector{X, 0}

		}
		tail = tail.add(unitVector)
		visitedPoints[point{tail.X, tail.Y}]++
	}
	return *head, *tail
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
			vectors = append(vectors, vector{X: 0, Y: -1 * value})
		} else if instruction[0] == "U" {
			vectors = append(vectors, vector{X: 0, Y: value})
		} else if instruction[0] == "L" {
			vectors = append(vectors, vector{X: -1 * value, Y: 0})
		} else if instruction[0] == "R" {
			vectors = append(vectors, vector{X: value, Y: 0})
		}
	}
	return vectors
}

func main() {
	filename := "input.txt"
	visitedPoints := map[point]int{point{0, 0}: 1}
	tail := point{0, 0}
	head := point{0, 0}
	vectors := readfileVectors(filename)
	for _, v := range vectors {
		head, tail = move(visitedPoints, v, &tail, &head)
	}
	fmt.Println(len(visitedPoints))
}
