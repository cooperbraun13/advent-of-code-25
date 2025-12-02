package main

// part 1

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

// countZeroHitsEndOnly simulates the dial and returns how many times
// it points at 0 after applying each rotation in order
func countZeroHitsEndOnly(rotations []Rotation) int {
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