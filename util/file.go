package util

import (
	"bufio"
	"io"
)

// ReadLines
// Read all lines from a reader and return them as a slice of strings.
// If an error occurs, it will be returned, unless it is io.EOF.
func ReadLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return nil, err
	}

	return lines, nil
}
