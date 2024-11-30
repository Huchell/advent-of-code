package aoc2015

import "testing"

func TestDayTwoPartOne(t *testing.T) {
	tests := []testCase{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	problem := DayTwo{}
	for _, testCase := range tests {
		result, err := problem.PartOne([]byte(testCase.string))
		if err != nil {
			t.Fatal(err)
		}

		if result != testCase.int {
			t.Fatalf("%s, %d != %d", testCase.string, result, testCase.int)
		}
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	tests := []testCase{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	problem := DayTwo{}
	for _, testCase := range tests {
		result, err := problem.PartTwo([]byte(testCase.string))
		if err != nil {
			t.Fatal(err)
		}

		if result != testCase.int {
			t.Fatalf("%s, %d != %d", testCase.string, result, testCase.int)
		}
	}
}
