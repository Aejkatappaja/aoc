package part1

import (
	"strings"
)

func Solve(file []byte) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	totalJoltage := 0

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}

		n := len(line)
		maxFrom := make([]int, n)
		maxFrom[n-1] = int(line[n-1] - '0')
		for i := n - 2; i >= 0; i-- {
			digit := int(line[i] - '0')
			maxFrom[i] = max(digit, maxFrom[i+1])
		}

		maxJoltage := 0
		for i := 0; i < n-1; i++ {
			tens := int(line[i] - '0')
			units := maxFrom[i+1]
			joltage := 10*tens + units
			maxJoltage = max(maxJoltage, joltage)
		}

		totalJoltage += maxJoltage
	}

	return totalJoltage, nil
}
