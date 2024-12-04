package aoc2024

import "testing"

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
	if err != nil {
		t.Fatal(err)
	}

	if result != 18 {
		t.Fatalf("%d != 18", result)
	}
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
	if err != nil {
		t.Fatal(err)
	}

	if result != 9 {
		t.Fatalf("%d != 9", result)
	}
}
