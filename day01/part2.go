package main

// part 2

// problem summary:
// - uses security method 0x434C49434B
// - instead of only checking at the end of each rotation, count every time the dial points to 0,
//   including intermediate clicks during a rotation
// - ex:
//    if R60 moves position one-by-one, multiple intermediate positions might land on 0
// - a rotation of say, R1000, can hit zero ten times, because each cycle of 100 positions
//   loops back to 0

// countZeroHitsAllClicks counts how many times the dial moves over 0,
// including during intermediate clicks, not just at the end
func countZeroHitsAllClicks(rotations []Rotation) int {
	const (
		startPosition = 50 // given by the problem
		dialSize = 100 // 0 -> 99, so 100 positions
	)

	position := startPosition
	zeroHits := 0

	// loop through each rotation instruction
	for _, rot := range rotations {
		switch rot.Dir {
		case 'L':
			// take rot.Steps single steps (key difference from part 1)
			for i := 0; i < rot.Steps; i++ {
				// move the dial one step left
				position--

				// if we go below 0, wrap around to 99
				if position < 0 {
					position += dialSize
				}

				// if after this single step we are at 0, count it
				if position == 0 {
					zeroHits++
				}
			}
		case 'R':
			// take rot.Steps single steps
			for i := 0; i < rot.Steps; i++ {
				// move the dial one step right
				position++

				// if we go past 99, wrap to 0
				if position >= dialSize {
					position -= dialSize
				}

				// check if we have arrived at 0
				if position == 0 {
					zeroHits++
				}
			}
		}
	}

	// total number of hits where dial == 0
	return zeroHits
}