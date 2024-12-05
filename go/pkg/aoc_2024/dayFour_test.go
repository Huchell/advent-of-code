package aoc2024

import (
	"huchell/aoc/pkg/assert"
	"testing"
)

func TestDayFourPartOne(t *testing.T) {
	input := []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	result, err := DayFour{}.PartOne(input)
	assert.ErrorIsNil(t, err)
	assert.IsEqual(t, result, 18)
}

func TestDayFourPartTwo(t *testing.T) {
	input := []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	result, err := DayFour{}.PartTwo(input)
	assert.ErrorIsNil(t, err)
	assert.IsEqual(t, result, 9)
}
