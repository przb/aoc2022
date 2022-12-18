package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}
	lines := make([]string, 0, stat.Size())

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		lines = append(lines, scn.Text())
	}
	return lines, nil
}

// ReadTokens returns a 2-d array for the file with the given filename
func ReadTokens(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}
	lines := make([][]string, 0, stat.Size())

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		line := strings.Split(scn.Text(), " ")
		lines = append(lines, line)
	}
	return lines, nil
}
