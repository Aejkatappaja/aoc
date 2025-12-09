package part2

import (
	"strings"
)

func Solve(file []byte) (int64, error) {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	var totalJoltage int64 = 0

	for _, line := range lines {
		if len(line) < 12 {
			continue
		}

		var joltage int64 = 0
		pos := 0

		for i := 0; i < 12; i++ {
			limite := len(line) - (12 - i)

			maxDigit := 0
			maxPos := pos
			for j := pos; j <= limite; j++ {
				digit := int(line[j] - '0')
				if digit > maxDigit {
					maxDigit = digit
					maxPos = j
				}
			}

			joltage = joltage*10 + int64(maxDigit)
			pos = maxPos + 1
		}

		totalJoltage += joltage
	}

	return totalJoltage, nil
}
