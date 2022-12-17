package util

import (
	"bufio"
	"os"
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
