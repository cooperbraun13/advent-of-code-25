package main

import (
	"bufio"
	"fmt"
	"os"
)

// advent of code 2025 day 3

// readLines loads all lines from the given file into a []string
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}

	// part 1
	part1 := sumMaxJoltageTwo(lines)
	fmt.Println("part 1:", part1)

	// part 2
	part2 := sumMaxJoltageTwelve(lines)
	fmt.Println("part 2:", part2)
}