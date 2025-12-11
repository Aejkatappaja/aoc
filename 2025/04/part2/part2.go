package part2

import (
	"strings"
)

const (
	rollOfPaper        = '@'
	removedRollOfPaper = '.'
)

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

	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}

	totalRemoved := 0

	for {
		toRemove := findAccessibleRolls(grid)

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			grid[pos[1]][pos[0]] = removedRollOfPaper
		}

		totalRemoved += len(toRemove)
	}

	return totalRemoved, nil
}

func findAccessibleRolls(grid [][]byte) [][2]int {
	accessible := [][2]int{}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == rollOfPaper {
				neighbors := 0

				for _, dir := range directions {
					neighbors += countNeighbors(x, y, dir, grid)
				}

				if neighbors < 4 {
					accessible = append(accessible, [2]int{x, y})
				}
			}
		}
	}
	return accessible
}

func countNeighbors(x, y int, dir []int, grid [][]byte) int {
	nx := x + dir[0]
	ny := y + dir[1]

	inBounds := nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[ny])

	if inBounds && grid[ny][nx] == rollOfPaper {
		return 1
	}
	return 0
}
