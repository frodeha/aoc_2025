package main

import (
	"aoc2025"
	"fmt"
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

	line := string(b)
	ranges := strings.Split(line, ",")

	part1(ranges)
	fmt.Printf("\n\n--- \n\n")
	part2(ranges)
}

func part1(ranges []string) {
	isSilly := func(s string) bool {
		if len(s) < 2 || len(s)%2 != 0 {
			return false
		}
		firstHalf := s[0 : len(s)/2]
		secondHalf := s[len(s)/2:]
		return firstHalf == secondHalf
	}

	sum := 0
	for _, r := range ranges {
		ends := strings.Split(r, "-")

		from, err := strconv.Atoi(ends[0])
		fatal(err)
		to, err := strconv.Atoi(ends[1])
		fatal(err)

		for ; from <= to; from++ {
			if isSilly(strconv.Itoa(from)) {
				debug("%d is silly\n", from)
				sum += from
			}
		}
	}

	fmt.Printf("Part 1: The sum of all the silly invalid ids is: %d\n", sum)
}

func part2(ranges []string) {
	isSilly := func(s string) bool {
		for slice := len(s) / 2; slice > 0; slice-- {
			if len(s)%slice != 0 {
				continue
			}

			part1 := s[:slice]
			silly := true
			for idx := slice; idx < len(s); idx += slice {
				part2 := s[idx : idx+slice]
				if part1 != part2 {
					silly = false
					break
				}
			}
			if silly {
				return true
			}
		}
		return false
	}

	sum := 0
	for _, r := range ranges {
		ends := strings.Split(r, "-")

		from, err := strconv.Atoi(ends[0])
		fatal(err)
		to, err := strconv.Atoi(ends[1])
		fatal(err)

		for ; from <= to; from++ {
			if isSilly(strconv.Itoa(from)) {
				debug("%d is silly\n", from)
				sum += from
			}
		}
	}

	fmt.Printf("Part 2: The sum of all the silly invalid ids is: %d\n", sum)
}
