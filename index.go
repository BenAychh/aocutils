package aocutils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// LoadFileAsStringArray loads text from a file with each line being an element of an array
func LoadFileAsStringArray(filename string) []string {
	file, err := os.Open(filename)
	handleError(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	handleError(scanner.Err())
	return lines
}

// LoadFileAsIntegerArray loads text from a file with each line being an element of an array
func LoadFileAsIntegerArray(filename string) []int {
	lines := LoadFileAsStringArray(filename)

	ints := []int{}

	for _, line := range lines {
		integer, err := strconv.Atoi(line)
		handleError(err)
		ints = append(ints, integer)
	}
	return ints
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
