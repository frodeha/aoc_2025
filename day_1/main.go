package main

import (
	"aoc2025"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	debug  = aoc2025.Debug
	fatal  = aoc2025.Fatal
	assert = aoc2025.Assert
)

func main() {
	b, err := os.ReadFile("full-input.txt")
	fatal(err)

	lines := strings.Split(string(b), "\n")
	part1(lines)
	debug("\n\n--\n\n")
	part2(lines)
}

func part1(lines []string) {
	var (
		zeroCount = 0
		pos       = 50
	)

	debug("L: %3s, pos: %d, zeros: %d\n", "XX", pos, zeroCount)
	for _, line := range lines {
		adjustment := adjustmentFromLine(line)
		pos = (pos + adjustment) % 100
		if pos < 0 {
			pos += 100
		}

		if pos == 0 {
			zeroCount += 1
		}

		assert(pos >= 0 && pos < 100, "position out of bounds, expected [0,100), was %d", pos)
		debug("L: %3s, pos: %2d, zeros: %d\n", line, pos, zeroCount)
	}

	fmt.Printf("The number of times the dial stopped on zero is: %d\n", zeroCount)
}

func part2(lines []string) {
	var (
		zeroCount = 0
		pos       = 50
	)

	debug("L: %5s, new pos: %3d, zeros: %d\n", "XX", pos, zeroCount)
	for _, line := range lines {
		adjustment := adjustmentFromLine(line)
		rotations := (adjustment - (adjustment % 100)) / 100
		if rotations < 0 {
			rotations *= -1
		}
		zeroCount += rotations

		newPos := (pos + adjustment) % 100
		if newPos < 0 {
			newPos += 100
		}

		if newPos == 0 {
			zeroCount++
		} else if pos != 0 {
			if adjustment < 0 && newPos > pos {
				zeroCount++
			}
			if adjustment > 0 && newPos < pos {
				zeroCount++
			}
		}
		pos = newPos

		debug("L: %5s, new pos: %3d, zeros: %d\n", line, pos, zeroCount)
		assert(pos >= 0 && pos < 100, "position out of bounds, expected [0,100), was %d", pos)
	}

	fmt.Printf("The number of times the dial passed zero is: %d\n", zeroCount)
}

func adjustmentFromLine(line string) int {
	direction := line[0]
	number, err := strconv.Atoi(string(line[1:]))
	fatal(err)

	var mod int
	switch direction {
	case 'L':
		mod = -1
	case 'R':
		mod = 1
	default:
		panic(fmt.Errorf("unexpected direction %c", direction))
	}

	return number * mod
}
