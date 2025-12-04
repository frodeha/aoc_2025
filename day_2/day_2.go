package day_2

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

func Challenge() day2 {
	return day2{}
}

type day2 struct{}

func (day2) Day() int {
	return 2
}

func (d day2) Part1() string {
	isSilly := func(s string) bool {
		if len(s) < 2 || len(s)%2 != 0 {
			return false
		}
		firstHalf := s[0 : len(s)/2]
		secondHalf := s[len(s)/2:]
		return firstHalf == secondHalf
	}

	sum := 0
	for _, r := range d.readInput() {
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

	return fmt.Sprintf("The sum of all the silly invalid ids is: %d", sum)
}

func (d day2) Part2() string {
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
	for _, r := range d.readInput() {
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

	return fmt.Sprintf("The sum of all the silly invalid ids is: %d", sum)
}

func (d day2) readInput() []string {
	b, err := os.ReadFile(aoc2025.InputFile(d.Day()))
	fatal(err)
	return strings.Split(string(b), ",")
}
