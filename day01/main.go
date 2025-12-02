package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// advent of code 2025 day 1

// Rotation represents a single instruction like "L68" or "R5"
type Rotation struct {
	Dir rune // 'L' or 'R'
	Steps int // number of clicks to move
}

func main() {
	// read all rotations from standard input
	rotations, err := readRotations()
	if err != nil {
		log.Fatalf("failed to read rotations: %v", err)
	}

	// count how many times we land on 0, part 1
	part1 := countZeroHitsEndOnly(rotations)

	// count how many times we pass by 0, part 2
	part2 := countZeroHitsAllClicks(rotations)

	// print answers for AoC website
	fmt.Println("part 1: ", part1)
	fmt.Println("part 2: ", part2)
}

// readRotations reads lines from stdin and parses them into []Rotation

// expected line formats:
//  L68
//  R5
//  L123

// it ignores blank lines
func readRotations() ([]Rotation, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var rotations []Rotation

	for scanner.Scan() {
		line := scanner.Text()
		// trim any leading/trailing whitespace just in case
		line = strings.TrimSpace(line)

		// skip any empty lines
		if line == "" {
			continue
		}

		// the first character should be either an 'L' or an 'R'
		dir := rune(line[0])
		if dir != 'L' && dir != 'R' {
			return nil, fmt.Errorf("invalid direction %q in line %q", dir, line)
		}

		// the rest of the string should be the number of steps, e.g. "68"
		stepsStr := strings.TrimSpace(line[1:])
		if stepsStr == "" {
			return nil, fmt.Errorf("missing step count in line %q", line)
		}

		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			return nil, fmt.Errorf("invalid step count %q in line %q: %w", stepsStr, line, err)
		}

		rotations = append(rotations, Rotation{
			Dir: dir,
			Steps: steps,
		})
	}

	// check if the scanner itself had an error while reading the stdin
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}