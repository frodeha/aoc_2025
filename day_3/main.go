package main

import (
	"aoc2025"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	debug = aoc2025.Debug
	fatal = aoc2025.Fatal
)

func main() {
	b, err := os.ReadFile("full-input.txt")
	fatal(err)

	lines := strings.Split(string(b), "\n")
	part1(lines)
	fmt.Printf("\n---\n\n")
	part2(lines)
}

func part2(lines []string) {
	var sum = 0
	for _, line := range lines {
		var (
			batteries = parse(line)
			picks     = 12

			maxJoltage = 0
			i          = 0
			leeway     = len(batteries) - (picks - 1)
		)

		for picks > 0 {
			max := 0
			newI := 0
			debug("checking range %d -> %d (%v) for pick %d\n", i, leeway, batteries[i:i+leeway], picks)
			for idx := i; idx < len(batteries) && idx < i+leeway; idx++ {
				joltage := batteries[idx]
				if joltage > max {
					max = joltage
					newI = idx
				}
			}
			debug("picked %d for pick %d\n", max, picks)

			picks--
			leeway -= newI - i
			i = newI + 1
			maxJoltage += max * int(math.Pow10(picks))
		}
		debug("Batteries: %s, Max joltage: %d\n", line, maxJoltage)
		sum += maxJoltage
	}
	fmt.Printf("Part 2: The total joltage output is: %d\n", sum)
}

func part1(lines []string) {
	var sum = 0
	for _, line := range lines {
		var (
			batteries  = parse(line)
			maxJoltage = 0
		)
		for idx, idxjoltage := range batteries {
			for _, jdxJoltage := range batteries[idx+1:] {
				newMax := (idxjoltage * 10) + jdxJoltage
				if newMax > maxJoltage {
					maxJoltage = newMax
				}
			}
		}

		debug("Batteries: %v, max joltage: %d\n", batteries, maxJoltage)
		sum += maxJoltage
	}

	fmt.Printf("Part 1: The total joltage output is: %d\n", sum)
}

func parse(line string) []int {
	batteries := make([]int, 0, len(line))
	for _, b := range strings.Split(line, "") {
		joltage, err := strconv.Atoi(b)
		fatal(err)
		batteries = append(batteries, joltage)
	}
	return batteries
}
