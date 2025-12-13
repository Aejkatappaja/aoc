package part1

import (
	"strconv"
	"strings"
)

func Solve(file []byte) (int, error) {
	totalFreshIngredients := 0
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	freshIngredientsRange := [][]int{}
	availableIngredients := []int{}

	// iterates on first file part before blank lines to find freshIngredientsIDs range list
	for _, line := range lines {
		if line == "" {
			break
		}

		rangeParts := strings.Split(line, "-")

		minID, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return 0, err
		}

		maxID, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return 0, err
		}

		idRange := []int{minID, maxID}

		freshIngredientsRange = append(freshIngredientsRange, idRange)
	}

	// iterates from blank line idx to the end to find freshIngredientsIDs and convert them to int
	startIdx := len(freshIngredientsRange) + 1
	for _, l := range lines[startIdx:] {
		ingredientID, err := strconv.Atoi(l)
		if err != nil {
			return 0, err
		}
		availableIngredients = append(availableIngredients, ingredientID)
	}

	// Check each available ingredient against the fresh ingredient ranges
	for _, ingredientID := range availableIngredients {
		isFresh := false
		for _, idRange := range freshIngredientsRange {
			minID := idRange[0]
			maxID := idRange[1]

			if ingredientID >= minID && ingredientID <= maxID {
				isFresh = true
				break
			}
		}
		if isFresh {
			totalFreshIngredients += 1
		}
	}

	return totalFreshIngredients, nil
}
