package day_5

import (
	"aoc2025"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	fatal = aoc2025.Fatal
	debug = aoc2025.Debug
)

func Challenge() day5 {
	return day5{}
}

type day5 struct{}

func (day5) Day() int {
	return 5
}

func (d day5) Part1() string {
	ranges, ids := d.readInput()
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.from && id <= r.to {
				count++
				break
			}
		}
	}

	return fmt.Sprintf("The number of ingredients that are fresh is: %d", count)
}

func (d day5) Part2() string {
	ranges, _ := d.readInput()

	slices.SortFunc(ranges, func(a, b _range) int {
		diff := a.from - b.from
		if diff == 0 {
			diff = (a.to - a.from) - (b.to - b.from)
		}
		return diff
	})

	debug("ranges is %v\n", ranges)

	prev := ranges[0]
	count := prev.to - prev.from + 1
	for _, r := range ranges[1:] {
		if r.from <= prev.to {
			r.from = prev.to + 1
		}

		if r.from <= r.to {
			count += r.to - r.from + 1
			prev = r
		}
	}

	return fmt.Sprintf("The number of ingredient ids can be considered to be fresh is: %d", count)
}

type _range struct {
	from int
	to   int
}

func (d day5) readInput() ([]_range, []int) {
	b, err := os.ReadFile(aoc2025.InputFile(d.Day()))
	fatal(err)

	var (
		ranges []_range
		ids    []int

		parsedRanges bool
	)

	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			parsedRanges = true
			continue
		}

		if parsedRanges {
			i, err := strconv.Atoi(line)
			fatal(err)
			ids = append(ids, i)
			continue
		}

		parts := strings.Split(line, "-")
		from, err := strconv.Atoi(parts[0])
		fatal(err)
		to, err := strconv.Atoi(parts[1])
		fatal(err)

		ranges = append(ranges, _range{from, to})
	}

	return ranges, ids
}
