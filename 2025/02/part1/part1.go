package part1

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(file []byte) (int, error) {
	ranges := strings.Split(strings.TrimSpace(string(file)), ",")

	sum := 0
	for _, r := range ranges {
		start, end, err := parseRange(r)
		if err != nil {
			return 0, err
		}

		for num := start; num <= end; num++ {
			if isRepeatedTwice(num) {
				sum += num
			}
		}
	}
	return sum, nil
}

func parseRange(r string) (int, int, error) {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range: %s", r)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func isRepeatedTwice(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}
