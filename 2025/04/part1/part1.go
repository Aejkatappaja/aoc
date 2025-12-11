package part1

import (
	"strings"
)

const rollOfPaper = '@'

var directions = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func Solve(file []byte) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	rollsOfPaperAccessibleByForklift := 0

	for y := range lines {
		for x := range len(lines[y]) {
			if lines[y][x] == rollOfPaper {
				surroundingRollsOfPaper := 0
				for _, dir := range directions {
					surroundingRollsOfPaper += countNeighbors(x, y, dir, lines)
				}
				if surroundingRollsOfPaper < 4 {
					rollsOfPaperAccessibleByForklift++
				}
			}
		}
	}
	return rollsOfPaperAccessibleByForklift, nil
}

func countNeighbors(x, y int, dir []int, lines []string) int {
	nx := x + dir[0]
	ny := y + dir[1]

	inBounds := nx >= 0 && ny >= 0 && ny < len(lines) && nx < len(lines[ny])

	if inBounds {
		if lines[ny][nx] == rollOfPaper {
			return 1
		}
	}
	return 0
}
