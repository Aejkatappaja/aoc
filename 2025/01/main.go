package main

import (
	_ "embed"

	"github.com/aejkatappaja/aoc_2025_day01/part1"
	"github.com/aejkatappaja/aoc_2025_day01/part2"
)

//go:embed part1/part1_input.txt
var inputBytes1 []byte

//go:embed part2/part2_input.txt
var inputBytes2 []byte

func main() {
	part1.Solve(inputBytes1)
	part2.Solve(inputBytes2)
}
