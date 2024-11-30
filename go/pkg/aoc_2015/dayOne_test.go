package aoc2015

import "testing"

type testCase struct {
	string
	int
}

func TestDayOnePartOne(t *testing.T) {
	tests := []testCase{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	problem := DayOne{}
	for _, testCase := range tests {
		result, err := problem.PartOne([]byte(testCase.string))
		if err != nil {
			t.Fatal(err)
		}
		result, ok := result.(int)
		if !ok {
			t.Fatal("Result is not int")
		}
		if result != testCase.int {
			t.Fatalf("%s, %d != %d", testCase.string, result, testCase.int)
		}
	}
}

func TestDayOnePartTwo(t *testing.T) {
	tests := []testCase{
		{")", 1},
		{"()())", 5},
	}

	problem := DayOne{}
	for _, testCase := range tests {
		result, err := problem.PartTwo([]byte(testCase.string))
		if err != nil {
			t.Fatal(err)
		}
		result, ok := result.(int)
		if !ok {
			t.Fatal("Result is not int")
		}
		if result != testCase.int {
			t.Fatalf("%s, %d != %d", testCase.string, result, testCase.int)
		}
	}
}
