package aoc2015

import (
	"testing"
)

func TestDayThreePartOne(t *testing.T) {
	tests := []testCase{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	problem := DayThree{}
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

func TestDayThreePartTwo(t *testing.T) {
	tests := []testCase{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	problem := DayThree{}
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
