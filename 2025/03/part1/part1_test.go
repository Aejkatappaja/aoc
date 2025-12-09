package part1

import (
	_ "embed"
	"testing"
)

//go:embed part1_test.txt
var inputBytesTest []byte

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		file     []byte
		expected int
	}{
		{
			name:     "example from puzzle",
			file:     inputBytesTest,
			expected: 357,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Solve(tt.file)
			if err != nil {
				t.Fatal(err)
			}

			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
