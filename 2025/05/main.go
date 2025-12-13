package main

import (
	_ "embed"
	"fmt"

	"github.com/aejkatappaja/aoc_2025_day05/part1"
	"github.com/aejkatappaja/aoc_2025_day05/part2"
)

//go:embed part1/part1_input.txt
var inputBytes1 []byte

//go:embed part2/part2_input.txt
var inputBytes2 []byte

func main() {
	solution1, err := part1.Solve(inputBytes1)
	if err != nil {
		panic(err)
	}
	fmt.Println(solution1)

	solution2, err := part2.Solve(inputBytes2)
	if err != nil {
		panic(err)
	}
	println(solution2)
}
