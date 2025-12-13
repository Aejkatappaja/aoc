package part2

import (
	"sort"
	"strconv"
	"strings"
)

func Solve(file []byte) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	type Range struct {
		min, max int
	}

	var ranges []Range

	// iterates on first file part before blank lines to find freshIngredientsIDs range list
	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		min, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		ranges = append(ranges, Range{min, max})
	}

	// Sort ranges by their minimum value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	total := 0
	prevMax := 0 // Track the maximum ID covered

	// Process each range in sorted order
	for _, r := range ranges {
		// Skip ranges completely covered by previous ranges
		if r.max <= prevMax {
			continue
		}

		// Calculate the start of the uncovered portion
		// If range starts after prevMax, use range.min
		// Otherwise, use prevMax+1 (first uncovered ID)
		start := max(r.min, prevMax+1)

		// Add count of uncovered IDs in this range
		total += r.max - start + 1

		// Update the maximum covered
		prevMax = max(r.max, prevMax)
	}

	return total, nil
}
