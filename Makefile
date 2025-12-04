build:
	@go build -ldflags "-X main.binary=aoc2025" -o aoc2025 cmd/main.go

day_1: build
	./aoc2025 day 1

day_2: build
	./aoc2025 day 2

day_3: build
	./aoc2025 day 3

day_4: build
	./aoc2025 day 4

all: build
	./aoc2025 all

.PHONY: build