package part2

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction string

const (
	dialRotateLeft    Direction = "L"
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

		direction := Direction(line[:1])
		dialClicks, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		delta := 1
		if direction == dialRotateLeft {
			delta = -1
		}

		for range dialClicks {
			dialPosition = ((dialPosition+delta)%dialSize + dialSize) % dialSize
			if dialPosition == 0 {
				zeroCount++
			}
		}

	}

	return zeroCount, nil
}
