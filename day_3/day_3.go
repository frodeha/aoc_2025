package day_3

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

func Challenge() day3 {
	return day3{}
}

type day3 struct{}

func (day3) Day() int {
	return 3
}

func (d day3) Part1() string {
	var sum = 0
	for _, line := range d.readInput() {
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

	return fmt.Sprintf("The total joltage output is: %d", sum)
}

func (d day3) Part2() string {
	var sum = 0
	for _, line := range d.readInput() {
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
	return fmt.Sprintf("The total joltage output is: %d", sum)
}

func (d day3) readInput() []string {
	b, err := os.ReadFile(aoc2025.InputFile(d.Day()))
	fatal(err)
	return strings.Split(string(b), "\n")
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
