package aoc2024

import (
	"testing"

	"huchell/aoc/pkg/assert"
)

func TestDayOnePartOne(t *testing.T) {
	input := []byte(`3   4
4   3
2   5
1   3
3   9
3   3
`)
	result, err := DayOne{}.PartOne(input)
	assert.TestErrorIsNil(t, err)
	assert.TestIsEqual(t, result, 11)
}

func TestDayOnePartTwo(t *testing.T) {
	input := []byte(`3   4
4   3
2   5
1   3
3   9
3   3
`)

	result, err := DayOne{}.PartTwo(input)
	assert.TestErrorIsNil(t, err)
	assert.TestIsEqual(t, result, 31)
}
