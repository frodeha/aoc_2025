package aoc2025

import (
	"fmt"
	"path"
)

var (
	DebugEnabled = false
	FullInput    = false
)

type Challenge interface {
	Day() int
	Part1() string
	Part2() string
}

func InputFile(day int) string {
	var (
		dayDir = fmt.Sprintf("day_%d", day)
		file   = "test.txt"
	)

	if FullInput {
		file = "full.txt"
	}

	return path.Join(dayDir, "input", file)
}

func Assert(b bool, s string, args ...interface{}) {
	if !b {
		Fatal(fmt.Errorf("%s %s", "[ASSERT]", fmt.Sprintf(s, args...)))
	}
}

func Debug(s string, args ...interface{}) {
	if DebugEnabled {
		fmt.Printf("%s %s", "[DEBUG]", fmt.Sprintf(s, args...))
	}
}

func Fatal(err error) {
	if err != nil {
		panic(fmt.Errorf("%s %s", "[FATAL]", err))
	}
}
