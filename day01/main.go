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

// problem summary:
// - have a dial with numbers 0 -> 99 arranged in a circle
// - it starts pointing at 50
// - each instruction is either:
//    L<number> -> rotate left (toward lower numbers)
//    R<number> -> rotate right (toward higher numbers)
// - the dial wraps around: going left from 0 goes to 99, and going right 99 goes to 0
// - need to count how many times the dial is exactly at 0 after applying each rotation, in sequence

// input:
// - each line in the input is one rotation, like: "L68" or "R5"

// output:
// - a single integer: the number of times the dial points at 0 after a rotation

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

	// count how many times we land on 0
	answer := countZeroHits(rotations)

	// print answer for AoC website
	fmt.Println(answer)
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

// countZeroHits simulates the dial and returns how many times
// it points at 0 after applying each rotation in order
func countZeroHits(rotations []Rotation) int {
	const (
		startPosition = 50 // given by the problem
		dialSize = 100 // 0 -> 99, so 100 positions
	)

	position := startPosition
	zeroHits := 0

	for _, rot := range rotations {
		// apply this rotation
		switch rot.Dir {
		case 'L':
			// left means "toward lower numbers"
			// example: from 11, L3 -> 8

			// do (position - steps) mod 100
			position = (position - rot.Steps) % dialSize

			// in Go, % can be negative, so have to normalize
			// if position is negative, wrap it back into 0 -> 99
			if position < 0 {
				position += dialSize
			}

		case 'R':
			// right means "toward higher numbers"
			// example: from 11, R3 -> 14

			// do (position + steps) mod 100
			position = (position + rot.Steps) % dialSize
		}

		// after applying this rotation, check if we landed on 0
		if position == 0 {
			zeroHits++
		}
	}
	
	return zeroHits
}