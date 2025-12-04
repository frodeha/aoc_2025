package day_4

import (
	"aoc2025"
	"fmt"
	"os"
	"strings"
)

var (
	fatal  = aoc2025.Fatal
	assert = aoc2025.Assert
)

func Challenge() day4 {
	return day4{}
}

type day4 struct{}

func (day4) Day() int {
	return 4
}

func (d day4) Part1() string {
	var (
		grid  = d.readInput()
		count = 0
	)
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			cell := grid.get(x, y)
			if cell != "@" {
				continue
			}

			adjacent := 0
			for xx := -1; xx < 2; xx++ {
				for yy := -1; yy < 2; yy++ {
					if xx == 0 && yy == 0 {
						continue
					}
					if grid.get(x+xx, y+yy) == "@" {
						adjacent++
					}
				}
			}
			if adjacent < 4 {
				count++
			}
		}
	}

	return fmt.Sprintf("The number of paper rolls that can be accessed by a forklift is: %d", count)
}

func (d day4) Part2() string {
	type tuple struct {
		x int
		y int
	}

	var (
		grid  = d.readInput()
		count = 0

		clearList []tuple
	)

	for {
		for _, tuple := range clearList {
			grid.clear(tuple.x, tuple.y)
		}
		clearList = nil

		for y := 0; y < grid.height; y++ {
			for x := 0; x < grid.width; x++ {
				cell := grid.get(x, y)
				if cell != "@" {
					continue
				}

				adjacent := 0
				for xx := -1; xx < 2; xx++ {
					for yy := -1; yy < 2; yy++ {
						if xx == 0 && yy == 0 {
							continue
						}
						if grid.get(x+xx, y+yy) == "@" {
							adjacent++
						}
					}
				}
				if adjacent < 4 {
					clearList = append(clearList, tuple{x, y})
				}
			}
		}

		count += len(clearList)
		if len(clearList) == 0 {
			break
		}
	}

	return fmt.Sprintf("The number of paper rolls that can be accessed by a forklift is: %d", count)
}

type grid struct {
	cells  [][]string
	width  int
	height int
}

func (g grid) get(x, y int) string {
	if x < 0 || x >= g.width {
		return ""
	}

	if y < 0 || y >= g.height {
		return ""
	}

	return g.cells[y][x]
}

func (g grid) clear(x, y int) {
	if x < 0 || x >= g.width {
		return
	}

	if y < 0 || y >= g.height {
		return
	}

	g.cells[y][x] = ""
}

func (d day4) readInput() grid {
	b, err := os.ReadFile(aoc2025.InputFile(d.Day()))
	fatal(err)

	var (
		cells [][]string
		cols  int
	)

	for _, row := range strings.Split(string(b), "\n") {
		if cols == 0 {
			cols = len(row)
		} else {
			assert(len(row) == cols, "uneven number of columns")
		}
		cells = append(cells, strings.Split(row, ""))
	}

	return grid{cells: cells, width: cols, height: len(cells)}
}
