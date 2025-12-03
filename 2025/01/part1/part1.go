package part1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction string

const (
	dialRotateLeft    Direction = "L"
	dialRotateRight   Direction = "R"
	dialStartPosition int       = 50
	dialSize          int       = 100
)

func Solve(file []byte) (int, error) {
	lines := strings.Split(string(file), "\n")

	password, err := countZeros(lines)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Printf("password: %v \n", password)
	return password, nil
}

func countZeros(lines []string) (zeroCount int, err error) {
	dialPosition := dialStartPosition

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dialRotationDirection := Direction(line[:1])

		dialClicks, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}

		dialPosition, err = applyDialRotation(dialRotationDirection, dialClicks, dialPosition)
		if err != nil {
			return 0, err
		}

		if dialPosition == 0 {
			zeroCount++
		}
	}

	return zeroCount, nil
}

func applyDialRotation(direction Direction, clicks, position int) (int, error) {
	switch direction {
	case dialRotateLeft:
		position -= clicks
	case dialRotateRight:
		position += clicks
	default:
		return 0, fmt.Errorf("unknown direction: %s", direction)
	}

	return ((position % dialSize) + dialSize) % dialSize, nil
}
