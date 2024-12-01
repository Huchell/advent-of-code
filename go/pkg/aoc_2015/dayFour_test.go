package aoc2015

import "testing"

func TestDayFourPartOne(t *testing.T) {
	tests := []testCase{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	problem := DayFour{}
	for _, test := range tests {
		result, err := problem.PartOne([]byte(test.string))
		if err != nil {
			t.Fatal(err)
		}
		if result != test.int {
			t.Fatalf("%s, %d != %d", test.string, result, test.int)
		}
	}
}
