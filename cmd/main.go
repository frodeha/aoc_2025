package main

import (
	"aoc2025"
	"aoc2025/day_1"
	"aoc2025/day_2"
	"aoc2025/day_3"
	"aoc2025/day_4"
	"strconv"

	"fmt"
	"os"
)

var binary string

var challenges = []aoc2025.Challenge{
	day_1.Challenge(),
	day_2.Challenge(),
	day_3.Challenge(),
	day_4.Challenge(),
}

func main() {
	if len(os.Args) < 2 {
		printUsage("Missing arguments")
		return
	}

	var day int
	for idx := 1; idx < len(os.Args); idx++ {
		arg := os.Args[idx]
		switch arg {
		case "--debug":
			aoc2025.DebugEnabled = true
		case "--full":
			aoc2025.FullInput = true
		case "all":
			day = -1
		case "day":
			if idx+1 == len(os.Args) {
				printUsage("Missing value for argument day")
				return
			}

			value := os.Args[idx+1]
			d, err := strconv.Atoi(value)
			if err != nil {
				printUsage("Invalid value %s for argument day: %s", value, err)
				return
			}
			day = d
			idx++
		default:
			printUsage("Invalid option %s", arg)
		}
	}

	if day == -1 {
		for _, c := range challenges {
			run(c)
		}
		return
	}

	for _, c := range challenges {
		if c.Day() == day {
			run(c)
			return
		}
	}
	fmt.Printf("Found no challenge for day %d\n", day)
}

func run(challenge aoc2025.Challenge) {
	output := challenge.Part1()
	fmt.Printf("[AoC 2025] Day %d - Part 1: %s\n", challenge.Day(), output)

	output = challenge.Part2()
	fmt.Printf("[AoC 2025] Day %d - Part 2: %s\n", challenge.Day(), output)
}

func printUsage(s string, args ...interface{}) {
	fmt.Printf(s, args...)
	fmt.Printf("\n\n")
	fmt.Printf("Usage: %s [all] [day=<day>]\n", binary)
}
