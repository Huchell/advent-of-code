package aoc2024

import (
	"huchell/aoc/pkg/assert"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	input := []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	result, err := DayTwo{}.PartOne(input)
	assert.TestErrorIsNil(t, err)
	assert.TestIsEqual(t, result, 2)
}

func TestDayTwoPartTwo(t *testing.T) {
	input := []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	result, err := DayTwo{}.PartTwo(input)
	assert.TestErrorIsNil(t, err)
	assert.TestIsEqual(t, result, 4)
}
